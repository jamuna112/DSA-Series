package sorting

import (
	"testing"
)

func TestColorSort(t *testing.T) {
	tests := []struct {
		name           string
		inputArr       []int
		expectedOutput []int
	}{
		{"random order", []int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{"single set", []int{2, 0, 1}, []int{0, 1, 2}},
		{"reverse order", []int{2, 2, 2, 1, 1, 1, 0, 0, 0}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
		{"single value", []int{0}, []int{0}},
		{"empty array", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputCopy := append([]int(nil), tt.inputArr...)

			SortColors(inputCopy)

			if !sliceEqual(inputCopy, tt.expectedOutput) {
				t.Errorf("test failed %s, expected %v but got %v", tt.name, tt.expectedOutput, inputCopy)
			}
		})

	}
}

func sliceEqual(nums1, nums2 []int) bool {

	if len(nums1) != len(nums2) {
		return false
	}

	for i, _ := range nums1 {
		if nums1[i] != nums2[i] {
			return false
		}
	}
	return true
}
