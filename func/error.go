package main

import (
	"fmt"
)

func SomeFunc() (int, error) {
	return 0, fmt.Errorf("error %d", 42)
}

func Error() {
	i, err := SomeFunc()
	// i := SomeFunc() // won't compile
	// i, _ := SomeFunc() // also ok
	if err == nil {
		fmt.Printf("Got %d.\n", i)
	} else {
		fmt.Printf("Error %v.\n", err)
	}
}
