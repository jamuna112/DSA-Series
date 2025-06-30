package sorting

import "fmt"

/*
	Bubble sort is a sorting algo. which starts from the first element
	of an array and compare with the second element. Whichever is greater
	element we swap them, it continues the process and
*/

func BubbleSort(arr []int) []int {

	for i := 0; i < len(arr)-1; i++ {

		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func PrintArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
	return arr
}
