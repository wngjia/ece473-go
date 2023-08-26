package main

import (
	"fmt"
	"log"
	"os"
)

func Defer() {
	file, err := os.Create("foo.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		file.Close()
		fmt.Println("File closed.")
	}() // the ending () actually calls the function

	for i := 0; i < 100; i++ {
		fmt.Fprintf(file, "%d\n", i)
	}
}
