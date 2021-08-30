package gulia

import (
	"errors"
	"runtime"
)

var queue = make(chan func())
var done = make(chan *struct{})

func keepOnSameThread() {
	// Force this function to keep running on the same thread
	runtime.LockOSThread()

	for f := range queue {
		f()
		done <- nil
	}
}

func do(f func()) {
	queue <- f
	<-done
}

func Open() {
	go keepOnSameThread()

	do(func() {
		jl_init()
	})
}

func Close() {
	do(func() {
		jl_atexit_hook(0)
	})
}

func EvalString(str string) (JuliaValue, error) {
	var v JuliaValue
	var e error

	do(func() {
		res := jl_eval_string(str)

		// Check for exceptions
		exc := jl_exception_occurred()
		if exc != nil {
			v, e = nil, errors.New(jl_typeof_str(exc))
			return
		}

		v, e = valueFromJulia(res)
	})

	return v, e
}

func GetFunction(name string) (JuliaFunction, error) {
	var f JuliaFunction
	var e error
	do(func() {
		f, e = functionFromJulia(jl_get_function(jl_main_module, name))
	})
	return f, e
}
