package search

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		index  int
	}{
		{"Element at the end", []int{15, 5, 20, 35, 2, 42, 67, 17}, 42, 5},
		{"Element at the beginning ", []int{21, 35, 2, 77, 170}, 21, 0},
		{"no target element", []int{15, 5, 20, 35, 2, 42, 67, 17}, 30, -1},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output := LinearSearch(tt.input, tt.target)
			if tt.index != output {
				t.Errorf("test failed, expected %d but got %d", tt.index, output)
			}
		})
	}
}
