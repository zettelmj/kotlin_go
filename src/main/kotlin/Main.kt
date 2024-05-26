package org.example

import java.lang.foreign.*
import java.lang.foreign.ValueLayout.*

fun readNullTerminatedString(ms: MemorySegment): String {
    val result = StringBuilder()
    for (i in 0..ms.byteSize()) {
        val c = ms.get(JAVA_BYTE, i)
        if (c.toInt() == 0x0) {
            break
        }
        result.append(c.toInt().toChar())
    }
    return result.toString()
}

fun main() {
    val C_POINTER = ADDRESS.withTargetLayout(MemoryLayout.sequenceLayout(Long.MAX_VALUE, JAVA_BYTE))

    Arena.ofConfined().use { arena ->
        val lookup = SymbolLookup.libraryLookup("/home/jens/code/untitled1/go-library/go-library.so", arena)

        val helloLib = Linker.nativeLinker().downcallHandle(
            lookup.find("helloLib").orElseThrow(),
            FunctionDescriptor.of(JAVA_INT, ADDRESS)
        )

        val str = arena.allocateUtf8String("lol1")

        val resultInt = helloLib.invokeExact(str) as Int

        //val result = readNullTerminatedString(resultPointer)

        val result = when(resultInt) {
            0 -> false
            1 -> true
            else -> throw Exception("Error $resultInt occurred")
        }

        println("Result '$result'")
    }
}