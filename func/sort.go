package main

import (
	"fmt"
	"sort"
)

func SortIndex(names []string) []int {
	indices := make([]int, 0)
	for i := range names {
		indices = append(indices, i)
	}
	sort.Slice(indices, func(l, r int) bool {
		lstr := names[indices[l]]
		rstr := names[indices[r]]
		return lstr < rstr
	})
	return indices
}

func Lambda() {
	names := []string{"Dave", "Bob", "Alice", "Clair"}
	for _, index := range SortIndex(names) {
		fmt.Printf("%s,", names[index])
	}
	fmt.Printf("names=%v\n", names)
}
