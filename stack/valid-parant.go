package stack

func IsValidBrackets(s string) bool {
	stack := []rune{}

	bracketMap := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {

		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else if len(stack) == 0 || stack[len(stack)-1] != bracketMap[ch] {
			return false
		} else {
			stack = stack[:len(stack)-1] //pop
		}
	}

	return len(stack) == 0
}
