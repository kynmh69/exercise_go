package main

import (
	"fmt"
	"math"
)

func main() {
	// Q1
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	tmp := l[0]
	for _, v := range l {
		tmp = int(math.Min(float64(tmp), float64(v)))
		// fmt.Println(tmp)
	}
	fmt.Printf("Min: %d\n", tmp)

	// Q2
	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	var sum int
	for _, v := range m {
		sum += v
	}

	fmt.Printf("The total is %d.\n", sum)
}
