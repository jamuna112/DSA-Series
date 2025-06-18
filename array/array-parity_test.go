package array

import "testing"

func TestSortArrayByParity(t *testing.T) {

	tests := []struct {
		name  string
		input []int
	}{
		{"single odd", []int{1}},
		{"single even", []int{2}},
		{"even only", []int{2, 4, 8, 10}},
		{"odd only", []int{1, 3, 5, 7}},
		{"mixed numbers", []int{3, 1, 2, 4}},
		{"empty array", []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numsCopy := append([]int(nil), tt.input...)
			output := SortArrayByParity(numsCopy)

			if !isNumsEven(output) {
				t.Errorf("sort by parity %v, even numbers are not before odds", numsCopy)
			}
		})
	}

}

// helper function
func isNumsEven(nums []int) bool {
	isOdd := false

	for _, num := range nums {
		if num%2 == 0 {
			if isOdd {
				return false
			}

		} else {
			isOdd = true
		}
	}
	return true
}
