package stringss

import (
	"strings"
)

func ReverseWords(s string) string {

	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	start := 0
	end := len(words) - 1

	for start < end {

		temp := words[start]
		words[start] = words[end]
		words[end] = temp

		//move the pointer
		start++
		end--
	}
	return strings.Join(words, " ")
}
