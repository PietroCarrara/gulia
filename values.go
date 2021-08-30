package gulia

import (
	"fmt"
	"unsafe"
)

func valueFromJulia(v jl_value_t) (jv JuliaValue, err error) {
	if jl_typeis(v, jl_float64_type) {
		jv, err = &juliaBaseValue{val: jl_unbox_float64(v)}, nil
	} else if jl_typeis(v, jl_string_type) {
		jv, err = &juliaBaseValue{val: jl_string_data(v)}, nil
	} else {
		jv, err = &juliaUnknownValue{}, nil
	}

	jv.setBase(v)
	return
}

func valueFromGo(v interface{}) (jl_value_t, func(), error) {
	switch t := v.(type) {
	case *juliaBaseValue:
		return t.base, nil, nil
	case *juliaUnknownValue:
		return t.base, nil, nil
	case float64:
		return jl_box_float64(t), nil, nil
	case string:
		ptr := cstring(t)
		return jl_cstr_to_string(ptr), func() { free(unsafe.Pointer(ptr)) }, nil
	default:
		return nil, nil, fmt.Errorf(`could not box for object "%v"`, t)
	}
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
	var t string
	do(func(){
		t = jl_typeof_str(j.base)
	})
	return t
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
