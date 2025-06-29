package stringss

import "testing"

func TestNumOf(t *testing.T) {
	tests := []struct {
		name          string
		inputJewels   string
		inputStones   string
		expectedCount int
	}{
		{"Basic case with matches", "aA", "aAAbbbb", 3},
		{"No  match", "z", "ZZ", 0},
		{"All  are ", "abc", "abcabc", 6},
		{"Empty ", "", "anything", 0},
		{"Empty ", "aA", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumJewelsInStones(tt.inputJewels, tt.inputStones)
			if result != tt.expectedCount {
				t.Errorf("Expected %d but got %d", tt.expectedCount, result)
			}
		})
	}
}
