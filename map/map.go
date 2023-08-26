// map/map.go
package main

import (
	"fmt"
)

func main() {
	months := make(map[string]int)
	months["Jan"] = 1
	months["Feb"] = 2
	fmt.Printf("Jan is month %d.\n", months["Jan"])

	fmt.Printf("Input a name: ")
	var name string
	fmt.Scanf("%s", &name)
	index, ok := months[name]
	if !ok {
		fmt.Printf("Unknown month %v.\n", name)
	} else {
		fmt.Printf("%v is month %d.\n", name, index)
	}
}
