package search

func RotatedSortedArray(input []int, target int) int {

	if len(input) == 0 {
		return -1
	}

	st := 0
	end := len(input) - 1

	for st <= end {

		mid := st + (end-st)/2

		if target == input[mid] {
			return mid
		}

		//left sorted
		if input[st] <= input[mid] {
			if input[st] <= target && target <= input[mid] {
				end = mid - 1
			} else {
				st = mid + 1
			}
		} else {
			if input[mid] <= target && target <= input[end] {
				st = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return -1
}
