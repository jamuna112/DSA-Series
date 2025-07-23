package stack

func NextGreaterElement1(nums1, nums2 []int) []int {

	stack := []int{}
	nextGreaterMap := make(map[int]int)

	//iterate through nums2, reverse order
	for _, num := range nums2 {

		for len(stack) > 0 && num > stack[len(stack)-1] {
			prev := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nextGreaterMap[prev] = num
		}
		stack = append(stack, num)
	}

	//for element left in stack there's no greater element
	for len(stack) > 0 {
		prev := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		nextGreaterMap[prev] = -1
	}

	//Build the result for nums1
	result := make([]int, len(nums1))

	for i, num := range nums1 {
		result[i] = nextGreaterMap[num]
	}

	return result
}
