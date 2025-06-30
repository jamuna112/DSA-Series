package stringss

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
		{"    multiple     spaces    here   ", "here spaces multiple"},
		{"one", "one"},
		{"    ", ""},
	}

	for _, tt := range tests {
		output := ReverseWords(tt.input)

		if !reflect.DeepEqual(output, tt.expectedOutput) {
			t.Errorf("failed, expected %v but got %v", tt.expectedOutput, output)
		}
	}
}
