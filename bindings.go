package gulia

/*
#cgo CFLAGS: -std=gnu99 -I'/usr/include/julia' -fPIC
#cgo LDFLAGS: -L'/usr/lib' -Wl,--export-dynamic -Wl,-rpath,'/usr/lib' -Wl,-rpath,'/usr/lib/julia' -ljulia
#include <julia.h>

jl_value_t *typeOf(jl_value_t *t) {
  return jl_typeof(t);
}

int _typeis(jl_value_t* obj, jl_datatype_t* type) {
	return jl_typeis(obj, type);
}
*/
import "C"
import "unsafe"

type jl_value_t *C.jl_value_t
type jl_datatype_t *C.jl_datatype_t

var jl_float64_type jl_datatype_t
var jl_partial_struct_type jl_datatype_t

func jl_init() {
	C.jl_init()

	jl_float64_type = C.jl_float64_type
	jl_partial_struct_type = C.jl_partial_struct_type
}

func jl_atexit_hook(status int) {
	C.jl_atexit_hook(C.int(status))
}

func jl_eval_string(str string) jl_value_t {
	ptr := C.CString(str)
	defer C.free(unsafe.Pointer(ptr))

	return C.jl_eval_string(ptr)
}

func jl_exception_occurred() jl_value_t {
	return C.jl_exception_occurred()
}

func jl_typeof_str(v jl_value_t) string {
	return C.GoString(C.jl_typeof_str(v))
}

func jl_typeis(v jl_value_t, t jl_datatype_t) bool {
	return C._typeis(v, t) != C.int(0)
}

func jl_unbox_float64(v jl_value_t) float64 {
	return float64(C.jl_unbox_float64(v))
}
