package search

import "testing"

func TestPeakInMountainArray(t *testing.T) {
	tests := []struct {
		name        string
		inputArr    []int
		expectedIdx int
	}{
		{"peak at right side", []int{1, 3, 5, 6, 4, 2}, 3},
		{"peak at mid", []int{0, 2, 4, 3, 1}, 2},
		{"peak at beginning", []int{0, 10, 5, 2}, 1},
		{"very small peak", []int{1, 2, 1}, 1},
		{"minimum valid mountain array", []int{0, 1, 0}, 1},
	}

	for _, tt := range tests {
		output := PeakInMountainArray(tt.inputArr)

		if output != tt.expectedIdx {
			t.Errorf("test failed for %s. Expected %d, but got %d\n", tt.name, tt.expectedIdx, output)
		}
	}
}
