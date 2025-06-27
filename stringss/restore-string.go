package stringss

func RestoreString(s string, indices []int) string {

	result := make([]rune, len(s))

	for i, val := range indices {
		result[val] = rune(s[i])
	}
	return string(result)
}
