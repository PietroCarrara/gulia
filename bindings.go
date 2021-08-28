package gulia

/*
#cgo CFLAGS: -std=gnu99 -I'/usr/include/julia' -fPIC
#cgo LDFLAGS: -L'/usr/lib' -Wl,--export-dynamic -Wl,-rpath,'/usr/lib' -Wl,-rpath,'/usr/lib/julia' -ljulia
#include <julia.h>

jl_value_t *typeOf(jl_value_t *t) {
  return jl_typeof(t);
}

char* _string_data(jl_value_t *s) {
	return ((char*)s + sizeof(void*));
}

int _typeis(jl_value_t* obj, jl_datatype_t* type) {
	return jl_typeis(obj, type);
}
*/
import "C"
import "unsafe"

type jl_value_t *C.jl_value_t
type jl_datatype_t *C.jl_datatype_t
type jl_module_t *C.jl_module_t
type jl_function_t *C.jl_function_t

var jl_float64_type jl_datatype_t
var jl_partial_struct_type jl_datatype_t
var jl_string_type jl_datatype_t
var jl_main_module jl_module_t

func free(ptr unsafe.Pointer) {
	C.free(ptr)
}

func cstring(str string) *C.char {
	return C.CString(str)
}

func jl_init() {
	C.jl_init()

	jl_float64_type = C.jl_float64_type
	jl_partial_struct_type = C.jl_partial_struct_type
	jl_string_type = C.jl_string_type
	jl_main_module = C.jl_main_module
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

func jl_box_float64(v float64) jl_value_t {
	return C.jl_box_float64(C.double(v))
}

func jl_get_function(mod jl_module_t, name string) jl_function_t {
	ptr := C.CString(name)
	defer C.free(unsafe.Pointer(ptr))

	return C.jl_get_function(mod, ptr)
}

func jl_call(function jl_function_t, values []jl_value_t) jl_value_t {
	return C.jl_call(function, (**C.jl_value_t)(&values[0]), C.int(len(values)))
}

func jl_cstr_to_string(str *C.char) jl_value_t {
	return C.jl_cstr_to_string(str)
}

func jl_string_data(v jl_value_t) string {
	return C.GoString(C._string_data(v))
}
