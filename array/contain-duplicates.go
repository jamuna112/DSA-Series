package array

func ContainDuplicates(nums []int) bool {

	seenDup := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if seenDup[nums[i]] {
			return true
		}
		seenDup[nums[i]] = true
	}

	return false
}
