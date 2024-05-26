package main

import (
	"C"
	"errors"
	"fmt"

	"github.com/open-policy-agent/opa/rego"
)
import "context"

//export helloLib
func helloLib(c *C.char) C.int {
	/*fmt.Println("Passed in:", C.GoString(c))
	newString := C.CString(C.GoString(c) + " Hallo")
	fmt.Println("Passed in:", C.GoString(newString))*/

	result, err := eval(C.GoString(c))
	if err != nil {
		fmt.Println(err)
		return C.int(2)
	}
	if result {
		return C.int(1)
	} else {
		return C.int(0)
	}

}

func eval(name string) (bool, error) {
	module := `
    package example.helloLib

    import rego.v1
    
    default result := false
    
    result if {
        input.name == "lol"
    }
    `

	ctx := context.TODO()

	query, err := rego.New(
		rego.Query("x = data.example.helloLib.result"),
		rego.Module("example.helloLib", module),
	).PrepareForEval(ctx)

	if err != nil {
		fmt.Println(err)
		return false, err
	}

	input := map[string]interface{}{
		"name": name,
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		fmt.Println(err)
		return false, err
	} else if len(results) == 0 {
		return false, errors.New("no result")
	} else if _, ok := results[0].Bindings["x"]; !ok {
		return false, errors.New("unexpected result type")
	} else {
		return results[0].Bindings["x"].(bool), nil
	}
}

func main() {
	println("This is just a library")
}
