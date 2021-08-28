package gulia

type JuliaFunction interface {
	Call(args ...interface{}) (JuliaValue, error)
}

func functionFromJulia(function jl_function_t) (JuliaFunction, error) {
	return &baseJuliaFunction{base: function}, nil
}

type baseJuliaFunction struct {
	base jl_function_t
}

func (j *baseJuliaFunction) Call(args ...interface{}) (JuliaValue, error) {
	boxes := make([]jl_value_t, 0, len(args))
	frees := make([]func(), 0, len(args))
	defer (func() {
		for _, free := range frees {
			free()
		}
	})()

	for _, arg := range args {
		box, free, err := valueFromGo(arg)
		if free != nil {
			frees = append(frees, free)
		}
		if err != nil {
			return nil, err
		}
		boxes = append(boxes, box)
	}

	return valueFromJulia(jl_call(j.base, boxes))
}
