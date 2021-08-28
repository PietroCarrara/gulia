package gulia

import (
	"errors"
)

func Open() {
	jl_init()
}

func Close() {
	jl_atexit_hook(0)
}

func EvalString(str string) (JuliaValue, error) {
	res := jl_eval_string(str)

	// Check for exceptions
	exc := jl_exception_occurred()
	if exc != nil {
		return nil, errors.New(jl_typeof_str(exc))
	}

	return fromJulia(res)
}
