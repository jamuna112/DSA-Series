package sorting

import "testing"

func TestSortByFequency(t *testing.T) {
	tests := []struct {
		inputString    string
		expectedString string
	}{
		{"tree", "eetr"},
		{"cccaaa", "cccaaa"},
		{"Aabb", "bbAa"},
	}

	for _, tt := range tests {
		t.Run(tt.inputString, func(t *testing.T) {
			result := FreqSorting(tt.inputString)

			if !hasSameFrequencies(result, tt.inputString) {
				t.Errorf("test failed! input: %s, expected a freq match with %s, got %s", tt.inputString, tt.expectedString, result)
			}
		})
	}
}

func hasSameFrequencies(a, b string) bool {
	count := func(s string) map[rune]int {
		m := make(map[rune]int)
		for _, ch := range s {
			m[ch]++
		}
		return m
	}
	return mapsEqual(count(a), count(b))
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}
