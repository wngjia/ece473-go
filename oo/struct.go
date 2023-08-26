package main

import (
	"fmt"
)

type Vertex struct {
	X, Y float64
}

func Struct() {
	v := Vertex{X: 1, Y: 2}
	fmt.Printf("%v, ", v)
	v.X, v.Y = 3, 4
	fmt.Printf("%+v\n", v)
}
