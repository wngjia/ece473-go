package main

import (
	"fmt"
)

type Circle2 struct {
	Vertex
	R float64
}

func (c *Circle2) Move(dx, dy float64) {
	c.Vertex.Move(dx, dy)
}

func (c Circle2) String() string {
	return fmt.Sprintf("Circle(%.3f,%.3f,r=%.3f)", c.X, c.Y, c.R)
}

func Embedding() {
	ms := []Movable{&Circle2{Vertex: Vertex{X: 3, Y: 4}, R: 5}}
	MoveAll(10, 20, ms)
	fmt.Printf("%v\n", ms)
}
