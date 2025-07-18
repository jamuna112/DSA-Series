package stack

func StockSpan(prices []int) []int {
	n := len(prices)
	span := make([]int, n)
	stack := []int{} //stack will hold index

	for i := 0; i < n; i++ {
		//pop until stack empty or we find a price greater than current
		for len(stack) > 0 && prices[stack[len(stack)-1]] <= prices[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			span[i] = i + 1 //all previous prices are <= current
		} else {
			span[i] = i - stack[len(stack)-1]
		}
		stack = append(stack, i)

	}
	return span
}
