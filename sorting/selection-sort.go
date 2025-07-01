package sorting

import "fmt"

/*

	selection sort is a simple sorting method where we divide the list into two parts.
	The sorted part and the unsorted part. At each step, we look through the unsorted part to find the
	smallest value, and then swap it with the first unsorted element. This
	grows the sorted part one item at a time.

	works by repeatedly selecting the minimum from the remaining elements.
	always takes the same amount of time, no matter how sorted the input already is.
	Easy to understand but not very efficient for large datasets.

	Time complexity: O(n^2)
*/

func SelectionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		sIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[sIndex] {
				sIndex = j
			}
		}
		arr[i], arr[sIndex] = arr[sIndex], arr[i]

	}
	return arr
}

func PrintSelectionArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	return arr
}
