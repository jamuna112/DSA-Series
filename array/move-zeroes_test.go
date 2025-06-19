package array

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"mix values case", []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{"all zeroes case", []int{0, 0, 0}, []int{0, 0, 0}},
		{"1 and 0's case", []int{0, 0, 1}, []int{1, 0, 0}},
		{"non-zero case", []int{1, 2, 3}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := append([]int(nil), tt.input...)
			output := MoveZeroes(inputCopy)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, output)
			}
		})
	}
}
