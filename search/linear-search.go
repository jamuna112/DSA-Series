package search

import "fmt"

func LinearSearch(input []int, target int) int {

	for i := 0; i < len(input); i++ {
		if input[i] == target {
			fmt.Printf("Target element %d found in index %d\n", target, i)
			return i
		}
	}
	return -1
}
