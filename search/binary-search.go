package search

func BinarySearch(input []int, target int) int {
	st := 0
	end := len(input) - 1

	for st <= end {

		mid := (st + end) / 2

		if target > input[mid] {
			st = mid + 1
		} else if target < input[mid] {
			end = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
