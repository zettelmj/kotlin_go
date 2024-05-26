from ctypes import *

lib = cdll.LoadLibrary("./go-library/go-library.so")
lib.helloLib.argtypes = [c_char_p]

lib.helloLib.restype=c_char_p
arg = create_string_buffer(b"Test")
print( lib.helloLib(arg))