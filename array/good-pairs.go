package array

/*
LEETCODE
1512. Number of Good Pairs
*/

func NumberOfGoodPairs(nums []int) int {

	var count int = 0

	freq := make(map[int]int)

	for _, num := range nums {
		count += freq[num]
		freq[num]++

	}
	return count
}
