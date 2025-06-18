package array

/* 
LEETCODE
905. Sort Array By Parity
*/

func SortArrayByParity(nums []int) []int {

	var leftPointer int = 0

	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			nums[leftPointer], nums[i] = nums[i], nums[leftPointer]
			leftPointer++
		}
	}
	return nums
}
