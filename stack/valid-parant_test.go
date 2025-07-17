package stack

import (
	"testing"
)

func TestValidParent(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			output := IsValidBrackets(tt.input)

			if output != tt.expected {
				t.Errorf("test failed for input %v, expected %v but got %v", tt.input, tt.expected, output)
			}
		})
	}

}
