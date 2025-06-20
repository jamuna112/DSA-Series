package array

import (
	"testing"
)

func TestNumberOfGoodPairs(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{1, 2, 3, 1, 1, 3},
			expected: 4,
		},
		{
			name:     "example 2",
			input:    []int{1, 1, 1, 1},
			expected: 6,
		},
		{
			name:     "example 3",
			input:    []int{1, 2, 3},
			expected: 0,
		},
		{
			name:     "empty array",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "single element",
			input:    []int{5},
			expected: 0,
		},
		{
			name:     "two duplicates",
			input:    []int{5, 5},
			expected: 1,
		},
		{
			name:     "mixed repeated",
			input:    []int{1, 2, 1, 2, 1},
			expected: 4, // pairs: (0,2), (0,4), (2,4), (1,3)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := NumberOfGoodPairs(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, got %d", tt.expected, result)
			}
		})
	}
}
