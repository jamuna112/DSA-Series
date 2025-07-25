package search

import "testing"

func TestRotatedBinarySearch(t *testing.T) {
	tests := []struct {
		name          string
		arr           []int
		target        int
		expectedIndex int
	}{
		{"Not rotated, target at start", []int{1, 2, 3, 4, 5}, 1, 0},
		{"Not rotated, target at end", []int{1, 2, 3, 4, 5}, 5, 4},
		{"Rotated, target at start", []int{4, 5, 6, 1, 2, 3}, 4, 0},
		{"Rotated, target at end", []int{4, 5, 6, 1, 2, 3}, 3, 5},
		{"Rotated, target in right half", []int{6, 7, 1, 2, 3, 4, 5}, 2, 3},
		{"Rotated, target in left half", []int{6, 7, 1, 2, 3, 4, 5}, 7, 1},
		{"Rotated, target is mid", []int{6, 7, 8, 1, 2, 3, 4, 5}, 1, 3},
		{"Rotated, not found", []int{6, 7, 8, 1, 2, 3, 4, 5}, 10, -1},
		{"Empty array", []int{}, 5, -1},
		{"Single element - found", []int{5}, 5, 0},
		{"Single element - not found", []int{5}, 3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := RotatedSortedArray(tt.arr, tt.target)

			if output != tt.expectedIndex {
				t.Errorf("Test failed for %s. Expected %d but got %d", tt.name, tt.expectedIndex, output)
			}
		})
	}
}
