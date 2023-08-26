package main

import (
	"fmt"
)

type Movable interface {
	Move(dx, dy float64)
}

func MoveAll(dx, dy float64, movables []Movable) {
	for _, m := range movables {
		m.Move(dx, dy)
	}
}

type Circle struct {
	X, Y, R float64
}

func (c *Circle) Move(dx, dy float64) {
	c.X += dx
	c.Y += dy
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle(%.3f,%.3f,r=%.3f)", c.X, c.Y, c.R)
}

/*
func (v Vertex) String() string {
	return fmt.Sprintf("Vertex(%.3f,%.3f)", v.X, v.Y)
}
*/

func Interface() {
	ms := []Movable{}
	ms = append(ms, &Vertex{X: 1, Y: 2})
	ms = append(ms, &Circle{X: 3, Y: 4, R: 5})
	MoveAll(10, 20, ms)
	fmt.Printf("%v\n", ms)
}
