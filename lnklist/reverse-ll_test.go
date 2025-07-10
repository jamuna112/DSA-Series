package lnklist

import (
	"testing"
)

func TestReverseLList(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{"multiple nodes", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"single node", []int{2}, []int{2}},
		{"empty list", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			inputList := buildList(tt.input)
			expectedList := buildList(tt.expected)

			result := ReverseList(inputList)

			if !isEqual(result, expectedList) {
				t.Errorf("ReverseList %v = %v, want %v", tt.input, listToSlice(result), tt.expected)
			}

		})
	}

}
