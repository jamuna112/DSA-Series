package stack

import (
	"reflect"
	"testing"
)

func TestNextGreaterElement(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{4, 5, 2, 25}, []int{5, 25, 25, -1}},
		{[]int{6, 8, 0, 1, 3}, []int{8, -1, 1, 3, -1}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := NextGreaterElement(tt.input)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("test failed, expected %v but got %v", tt.expected, got)
			}
		})
	}
}
