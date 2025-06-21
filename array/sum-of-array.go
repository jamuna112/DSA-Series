package array

/*
lEETCODE
1480. Running Sum of 1d Array
*/

func sumOf1dArray(nums []int) []int {

	result := []int{}
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		result = append(result, sum)
	}
	return result
}
