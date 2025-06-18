package array

import (
	"reflect"
	"testing"
)

func TestShuffleArray(t *testing.T) {

	tests := []struct {
		input    []int
		name     string
		num      int
		expected []int
	}{
		{
			name:     "basic case",
			input:    []int{2, 5, 1, 3, 4, 7},
			expected: []int{2, 3, 5, 4, 1, 7},
			num:      3,
		},
		{
			name:     "single pair",
			input:    []int{2, 4},
			expected: []int{2, 4},
			num:      1,
		},
		{
			name:     "empty input",
			input:    []int{},
			expected: []int{},
			num:      0,
		},
		{
			name:     "pairs",
			input:    []int{1, 2, 3, 4},
			expected: []int{1, 3, 2, 4},
			num:      2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ShuffleArray(tt.num, tt.input)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("given %v does not match with %v", tt.expected, output)
			}

		})

	}
}
