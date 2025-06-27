package stringss

import "testing"

func TestRestoreString(t *testing.T) {

	tests := []struct {
		name           string
		inputString    string
		inputIndices   []int
		expectedOutput string
	}{
		{"Basic shuffle", "codeleet", []int{4, 5, 6, 7, 0, 2, 1, 3}, "leetcode"},
		{"No shuffle needed", "abc", []int{0, 1, 2}, "abc"},
		{"custom order", "aiohn", []int{3, 1, 4, 2, 0}, "nihao"},
		{"empty input", "", []int{}, ""},
	}

	for _, tt := range tests {
		result := RestoreString(tt.inputString, tt.inputIndices)

		if result != tt.expectedOutput {
			t.Errorf("Test failed, expected %v but got %v", tt.expectedOutput, result)
		}
	}
}
