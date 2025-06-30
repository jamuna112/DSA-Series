package sorting

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		name           string
		inputArr       []int
		expectedOutput []int
	}{
		{"random order", []int{5, 2, 1, 3, 4}, []int{1, 2, 3, 4, 5}},
		{"sorted order", []int{11, 12, 13, 14, 15}, []int{11, 12, 13, 14, 15}},
		{"reverse order", []int{20, 19, 18, 17, 16}, []int{16, 17, 18, 19, 20}},
		{"single value", []int{9}, []int{9}},
		{"empty array", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := BubbleSort(tt.inputArr)

			if !reflect.DeepEqual(output, tt.expectedOutput) {
				t.Errorf("%s test failed! expected %v but got %v", tt.name, tt.expectedOutput, output)
			}
		})

	}
}

/*
 test for PrintArray isn't very helpful since it only prints and return the same input
 writing unit test for it doesn't add value, so skipping that function
*/
