package search

import (
	"reflect"
	"testing"
)

func TestIndexOfAllOccurences(t *testing.T) {
	tests := []struct {
		name            string
		input           []int
		target          int
		expectedIndices []int
	}{
		{"multiple matches element", []int{1, 2, 3, 2, 4, 7, 2, 11, 2}, 2, []int{1, 3, 6, 8}},
		{"single match element", []int{10, 20, 30, 40}, 30, []int{2}},
		{"no match", []int{20, 30, 40, 50}, 60, []int{}},
		{"all matches element", []int{10, 10, 10, 10}, 10, []int{0, 1, 2, 3}},
		{"negative element matches", []int{-1, -5, -1, 3, -1}, -1, []int{0, 2, 4}},
		{"empty slice", []int{}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			output := IndexOfAllOccurences(tt.input, tt.target)

			if !reflect.DeepEqual(tt.expectedIndices, output) {
				t.Errorf("test failed, expected %v but got %v", tt.expectedIndices, output)
			}
		})
	}
}
