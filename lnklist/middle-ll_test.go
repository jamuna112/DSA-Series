package lnklist

import "testing"

func TestMiddleNode(t *testing.T) {

	tests := []struct {
		name         string
		inputSlice   []int
		expectedSlow int
	}{
		{"even nodes", []int{1, 2, 3, 4, 5, 6}, 4},
		{"odd nodes", []int{1, 2, 3, 4, 5}, 3},
		{"single number", []int{65}, 65},
		{"empty list", []int{}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			head := buildList(tt.inputSlice)

			got := MiddleNode(head)

			if got == nil {
				if tt.inputSlice != nil && len(tt.inputSlice) > 0 {
					t.Errorf("test failed, expected middle node is %v, but got %v", tt.inputSlice, tt.expectedSlow)
				}
			} else if got.Val != tt.expectedSlow {
				t.Errorf("middle node %v = %d, want %d", tt.inputSlice, got.Val, tt.expectedSlow)
			}

		})
	}
}
