package search

import "fmt"

/*
Find Minimum and Maximum
Problem:
Without using built-in functions, find the smallest and largest elements in an array.
*/

func MinAndMax(input []int) (int, int, error) {

	if len(input) == 0 {

		return 0, 0, fmt.Errorf("slice is empty")
	}

	min, max := input[0], input[0]

	for i := 0; i < len(input); i++ {
		if input[i] > max {
			max = input[i]
		}
		if input[i] < min {
			min = input[i]
		}
	}

	return min, max, nil
}
