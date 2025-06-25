package stringss

import (
	"unicode"
)

func isAlphaNum(str byte) bool {
	return unicode.IsLetter(rune(str)) || unicode.IsNumber(rune(str))
}

func isLower(str byte) byte {
	return byte(unicode.ToLower(rune(str)))
}

func IsPalindrome(str string) bool {
	st := 0
	end := len(str) - 1

	for st < end {
		if !isAlphaNum(str[st]) {
			st++
			continue

		}

		if !isAlphaNum(str[end]) {
			end--
			continue
		}

		if isLower(str[st]) != isLower(str[end]) {
			return false
		}
		st++
		end--
	}
	return true
}
