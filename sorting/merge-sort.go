package sorting

/*
Merge Sort is a sorting algorithm that works by dividing the input array into smaller halves, sorting those halves independently,
and then merging them back together in order.

The idea is simple:

If the array has only one item, it's already sorted.

If it's longer, split it into two parts, recursively sort each part, and then combine the results in sorted order.

The key strength of Merge Sort is that it keeps the sorting process well-organized and consistent, even for large datasets.
 It's predictable, has a guaranteed time complexity of O(n log n), and doesn’t get slower in worst-case scenarios like some other algorithms.

*/

func MergeSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(arr, left, right)

}

func merge(arr []int, left, right []int) []int {

	i, j := 0, 0
	result := []int{}

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
