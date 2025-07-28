package search

func PeakInMountainArray(input []int) int {

	st := 0
	end := len(input) - 1

	for st < end {

		mid := st + (end-st)/2

		if input[mid] < input[mid+1] {
			st = mid + 1
		} else {
			end = mid
		}
	}
	return st
}
