package array

/*
LEETCODE
1470. Shuffle the Array

*/

func ShuffleArray(n int, nums []int) []int {

	var result = make([]int, 2*n)

	for i := 0; i < n; i++ {
		result[2*i] = nums[i]
		result[2*i+1] = nums[i+n]
	}
	return result
}
