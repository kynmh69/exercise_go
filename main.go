package main

import (
	"exe_lib/lib"
	"fmt"
)

type Human interface {
	Say()
}

func main() {
	vertex := lib.New(1, 2, 3)
	fmt.Println(vertex)
}
