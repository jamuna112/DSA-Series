package stringss

func AnagramChecker(str1, str2 string) bool {

	wordMap1 := map[rune]int{}
	wordMap2 := map[rune]int{}

	runeChr1 := []rune(str1)
	runeChr2 := []rune(str2)

	for _, ch := range runeChr1 {
		wordMap1[ch]++
	}

	for _, chr := range runeChr2 {
		wordMap2[chr]++
	}

	if len(wordMap1) != len(wordMap2) {
		return false
	}

	for val := range wordMap1 {

		value2, okay := wordMap2[val]

		if !okay || value2 != wordMap1[val] {
			return false
		}
	}

	return true

}
