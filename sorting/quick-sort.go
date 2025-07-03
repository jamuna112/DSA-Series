package sorting

/*

QUICK SORT:
  - Pick a pivot - book for example
  - Divide into smaller and bigger groups - put thinner books to the left and thicker to the right
  - Do the same for the left pile and right pile
  - Combine the results

*/

func partition(arr []int, st, end int) int {

	idx := st - 1
	pivot := arr[end]

	for j := st; j < end; j++ {
		if arr[j] < pivot {
			idx++
			arr[j], arr[idx] = arr[idx], arr[j]
		}
	}
	idx++
	arr[idx], arr[end] = arr[end], arr[idx]

	return idx
}

func QuickSort(arr []int, st, end int) {

	if st < end {
		pIdx := partition(arr, st, end)
		QuickSort(arr, st, pIdx-1)
		QuickSort(arr, pIdx+1, end)
	}

}
