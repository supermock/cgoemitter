package cgoemitter

//#include <stdlib.h>
import "C"
import "unsafe"

//Arguments | Custom arguments received from C language
type Arguments []unsafe.Pointer

//free | Clears the C language pointers
func (args Arguments) free() {
	for _, arg := range args {
		C.free(arg)
	}
}

//Count | Returns the number of arguments
func (args Arguments) Count() int {
	return len(args)
}

//Arg | Get the index argument
func (args Arguments) Arg(index int) unsafe.Pointer {
	return args[index]
}

//Int | Converts the received pointer to integer type
func (args Arguments) Int(index int) int {
	return int(*(*C.int)(args[index]))
}

//Float | Converts the received pointer to the float32 type
func (args Arguments) Float(index int) float32 {
	return float32(*(*C.float)(args[index]))
}

//Double | Converts the received pointer to the float64 type
func (args Arguments) Double(index int) float64 {
	return float64(*(*C.double)(args[index]))
}

//String | Converts the received pointer to string type
func (args Arguments) String(index int) string {
	return C.GoString((*C.char)(args[index]))
}
