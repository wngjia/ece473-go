package main

import "fmt"

func Sum(a ...int) int {
	sum := 0
	for _, i := range a {
		sum += i
	}
	return sum
}

func Variadic() {
	fmt.Printf("sum(1,3,4)=%d\n", Sum(1, 3, 4))
	fmt.Printf("sum(1,2,3,4,5)=%d\n", Sum(1, 2, 3, 4, 5))
}
