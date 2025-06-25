package stringss

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {

	tests := []struct {
		input    []byte
		expected []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{}, []byte{}},
		{[]byte{'o', 'r', 'a', 'n', 'g', 'e'}, []byte{'e', 'g', 'n', 'a', 'r', 'o'}},
	}

	for _, tt := range tests {
		//copy the input slice
		inputCopy := make([]byte, len(tt.input))
		copy(inputCopy, tt.input)
		output := ReverseString(inputCopy)

		if !reflect.DeepEqual(output, tt.expected) {
			t.Errorf("Failed, expecting %v but got %v", tt.expected, output)
		}

	}
}
