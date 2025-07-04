package sorting

import (
	"sort"
	"strings"
)

func FreqSorting(str string) string {

	freqMap := map[rune]int{}
	var builder strings.Builder

	for _, ch := range str {
		freqMap[ch]++
	}

	chars := make([]rune, 0, len(freqMap))

	for ch := range freqMap {
		chars = append(chars, ch)
	}

	sort.Slice(chars, func(i, j int) bool {
		return freqMap[chars[i]] > freqMap[chars[j]]
	})

	for _, ch := range chars {
		builder.WriteString(strings.Repeat(string(ch), freqMap[ch]))

	}

	return builder.String()
}
