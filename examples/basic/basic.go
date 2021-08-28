package main

import (
	"fmt"

	"github.com/PietroCarrara/gulia"
)

func main() {
	gulia.Open()
	defer gulia.Close()

	plot, _ := gulia.EvalString(`
using Plots

plot(sin, 0:10)
	`)

	savefig, _ := gulia.GetFunction("savefig")
	fmt.Println(savefig.Call(plot, "plot.png"))
}
