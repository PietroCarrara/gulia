package main

import (
	"fmt"

	"github.com/PietroCarrara/gulia"
)

func main() {
	gulia.Open()
	defer gulia.Close()

	gulia.EvalString(`f(a) = "The value is: " * string(a)`)
	f, _ := gulia.GetFunction("f")
	val, _ := f.Call(12.34)
	maybeStr, _ := val.GetValue()
	str, _ := maybeStr.(string)

	fmt.Println(str)
}
