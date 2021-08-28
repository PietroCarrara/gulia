package main

import (
	"fmt"

	"github.com/PietroCarrara/gulia"
)

func main() {
	gulia.Open()
	defer gulia.Close()

	val, _ := gulia.EvalString(`
f(x) = sin(x) + x^2

f(12.34)
	`)

	fmt.Println(val.GetValue())
}
