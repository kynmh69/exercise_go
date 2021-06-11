package main

import "fmt"

func main() {
	n := 100
	p := &n
	fmt.Println(p)
	fmt.Println(*p)
}
