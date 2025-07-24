package search

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		arr           []int
		target        int
		expectedIndex int
	}{
		{"target before mid", []int{5, 6, 7, 8, 9, 10, 11, 12, 13}, 5, 0},
		{"target after mid", []int{5, 6, 7, 8, 9, 10, 11, 12, 13}, 13, 8},
		{"target mid", []int{5, 6, 7, 8, 9, 10, 11, 12, 13}, 9, 4},
		{"no matched target ", []int{5, 6, 7, 8, 9, 10, 11, 12, 13}, 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := BinarySearch(tt.arr, tt.target)

			if output != tt.expectedIndex {
				t.Errorf("Test failed for %s. Expected %d but got %d", tt.name, tt.expectedIndex, output)
			}
		})
	}
}
