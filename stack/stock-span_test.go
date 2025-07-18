package stack

import (
	"reflect"
	"testing"
)

func TestCalculateSpan(t *testing.T) {
	tests := []struct {
		name     string
		price    []int
		expected []int
	}{
		{"classic example", []int{100, 80, 60, 70, 60, 75, 85}, []int{1, 1, 1, 2, 1, 4, 6}},
		{"Single day", []int{42}, []int{1}},
		{"increasing", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"decreasing", []int{5, 4, 3, 2, 1}, []int{1, 1, 1, 1, 1}},
		{"mixed values", []int{10, 10, 10}, []int{1, 2, 3}},
		{"mixed-and-fall", []int{10, 4, 5, 90, 120, 80}, []int{1, 1, 2, 4, 5, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StockSpan(tt.price)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("test failed for %s, expected %v but got %v", tt.name, tt.expected, got)
			}
		})
	}
}
