package gulia

import "fmt"

func fromJulia(v jl_value_t) (jv JuliaValue, err error) {
	if jl_typeis(v, jl_float64_type) {
		jv, err = &juliaBaseValue{val: jl_unbox_float64(v)}, nil
	} else {
		jv, err = &juliaUnknownValue{}, nil
	}

	jv.setBase(v)
	return
}

type JuliaValue interface {
	GetType() string
	GetValue() (interface{}, error)

	setBase(jl_value_t)
}

type juliaBaseValue struct {
	base jl_value_t
	val  interface{}
}

type juliaUnknownValue struct {
	juliaBaseValue
}

func (j *juliaBaseValue) GetType() string {
	return jl_typeof_str(j.base)
}

func (j *juliaBaseValue) GetValue() (interface{}, error) {
	return j.val, nil
}

func (j *juliaBaseValue) setBase(base jl_value_t) {
	if j.base != nil {
		panic("setting the value of a object that already has a value")
	}
	j.base = base
}

func (j *juliaUnknownValue) GetValue() (interface{}, error) {
	return nil, fmt.Errorf("objects of type \"%s\" are not supported", j.GetType())
}
