package sorting

import "fmt"

func InsertionSort(arr []int) []int {

	for i := 1; i < len(arr); i++ { // 12, 11, 13, 5, 6

		current := arr[i] // 11
		prev := i - 1     // 12

		for prev >= 0 && arr[prev] > current { //  12 > 0 &&  12 > 11

			arr[prev+1] = arr[prev]
			prev--
		}
		arr[prev+1] = current

	}

	return arr
}

func PrintInsertionArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	return arr
}
