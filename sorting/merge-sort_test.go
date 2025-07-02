package sorting

import (
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name        string
		inputArr    []int
		expectedArr []int
	}{
		{"random order", []int{64, 25, 11, 22, 12}, []int{11, 12, 22, 25, 64}},
		{"reverse order", []int{90, 80, 70, 60, 50, 40, 30, 20, 10}, []int{10, 20, 30, 40, 50, 60, 70, 80, 90}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"single number", []int{9}, []int{9}},
		{"empty array set", []int{}, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := MergeSort(tt.inputArr)

			if !reflect.DeepEqual(output, tt.expectedArr) {
				t.Errorf("%s test failed, expected %v but got %v", tt.name, tt.expectedArr, output)
			}
		})
	}
}
