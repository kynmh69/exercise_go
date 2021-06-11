package lib

import "fmt"

func sub3() {
	n := 100
	p := &n
	fmt.Println(p)
	fmt.Println(*p)
}
