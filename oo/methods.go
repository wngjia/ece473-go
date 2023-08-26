package main

import (
	"fmt"
	"math"
)

func (v *Vertex) Move(dx, dy float64) {
	v.X += dx
	v.Y += dy
}

func (v Vertex) Norm() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Methods() {
	v := Vertex{X: 1, Y: 2}
	v.Move(1, 2)
	fmt.Printf("%+v, norm=%.3f\n", v, v.Norm())
}
