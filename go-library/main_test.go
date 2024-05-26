package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRules(t *testing.T) {
	result, err := eval("lol")
	fmt.Printf("%+v", result)
	assert.Equal(t, true, result)
	assert.Nil(t, err)
}
