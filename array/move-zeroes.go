package array

func MoveZeroes(nums []int) []int {
	var p int = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[p] = nums[i]
			p++
		}

	}
	for p < len(nums) {
		nums[p] = 0
		p++
	}
	return nums
}
