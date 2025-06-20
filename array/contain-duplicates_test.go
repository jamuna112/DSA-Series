package array

import (
	"reflect"
	"testing"
)

func TestContainDuplicates(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"unique values", []int{1, 2, 3, 4}, false},
		{"duplicate values", []int{1, 2, 3, 1}, true},
		{"two same numbers", []int{5, 5}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := ContainDuplicates(tt.input)

			if !reflect.DeepEqual(output, tt.expected) {
				t.Errorf("expected %v got %v", tt.expected, output)
			}
		})
	}
}
