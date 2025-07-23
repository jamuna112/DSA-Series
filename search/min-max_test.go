package search

import "testing"

func TestMinAndMax(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		expectedMin   int
		expectedMax   int
		expectedError bool
	}{
		{"With all positive elements", []int{90, 1, 40, 34, 21, 50, 27}, 1, 90, false},
		{"mix of positive and negative elements", []int{-20, 12, -3, 50, 30}, -20, 50, false},
		{"empty slice", []int{}, 0, 0, true},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			min, max, err := MinAndMax(tt.input)

			if tt.expectedError {
				if err == nil {
					t.Errorf("Expected error but got nil")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}

			if min != tt.expectedMin {
				t.Errorf("Test failed, expected min is %d but got %d", tt.expectedMin, min)
			}

			if max != tt.expectedMax {
				t.Errorf("Test failed, expected max id %d but got max %d", tt.expectedMax, max)
			}
		})
	}
}
