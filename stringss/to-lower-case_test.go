package stringss

import (
	"reflect"
	"testing"
)

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"HELLO", "hello"},
		{"Kyle", "kyle"},
		{"jefferson", "jefferson"},
		{"", ""},
	}

	for _, tt := range tests {
		output := toLowerCase(tt.input)

		if !reflect.DeepEqual(output, tt.expectedOutput) {
			t.Errorf("Failed, expected %v but got %v", tt.expectedOutput, output)
		}
	}
}
