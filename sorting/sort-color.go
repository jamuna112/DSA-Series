package sorting

func SortColors(nums []int) {

	low := 0
	mid := 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			mid++
			low++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
}
