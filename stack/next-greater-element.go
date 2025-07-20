package stack

func NextGreaterElement(elements []int) []int {

	n := len(elements)
	ans := make([]int, n)
	st := []int{}

	for i := n - 1; i >= 0; i-- {
		for len(st) > 0 && elements[st[len(st)-1]] <= elements[i] {
			st = st[:len(st)-1]
		}

		//if stack is empty, no greate element
		if len(st) == 0 {
			ans[i] = -1
		} else {
			ans[i] = elements[st[len(st)-1]]
		}
		//push current index to stack
		st = append(st, i)
	}
	return ans
}
