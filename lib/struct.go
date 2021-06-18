package lib

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y int
}

func (v Vertex) Area() int {
	return v.x * v.y
}

func (v *Vertex) Scale(i int) {
	v.x *= i
	v.y *= i
}

func New(x, y, z int) *Vertex3D {
	return &Vertex3D{Vertex{x, y}, z}
}

type Vertex3D struct {
	Vertex
	z int
}

func (v Vertex3D) Area3D() int {
	return v.x * v.y * v.z

}

func (v *Vertex3D) Scale3D(i int) {
	v.x *= i
	v.y *= i
	v.z *= i
}

func sub1() {
	// Q1
	i := 1.11
	fmt.Println(int(i))

	// Q2
	// [5, 6]

	// Q3
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	fmt.Printf("%T %v", m, m)
}

func sub2() {
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

func sub3() {
	n := 100
	p := &n
	fmt.Println(p)
	fmt.Println(*p)
}
