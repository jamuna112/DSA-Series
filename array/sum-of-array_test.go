package array

import (
	"reflect"
	"testing"
)

func TestSumOf1dArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"simple case", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"random case", []int{3, 1, 4, 1, 5}, []int{3, 4, 8, 9, 14}},
		{"all zeroes", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{"empty case", []int{}, []int{}},
	}

	for _, tt := range tests {
		result := sumOf1dArray(append([]int{}, tt.input...))

		if !reflect.DeepEqual(result, tt.expected) {
			t.Errorf("expected %v but got %v", tt.expected, result)
		}
	}
}
