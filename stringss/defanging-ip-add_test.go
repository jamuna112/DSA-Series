package stringss

import (
	"testing"
)

func TestDefangIPaddr(t *testing.T) {

	tests := []struct {
		input          string
		expectedOutput string
	}{
		{"255.100.50.0", "255[.]100[.]50[.]0"},
		{"1.1.1.1", "1[.]1[.]1[.]1"},
		{"", ""},
		{"192.168.0.1", "192[.]168[.]0[.]1"},
	}

	for _, tt := range tests {
		output := DefangIPaddr(tt.input)

		if output != tt.expectedOutput {
			t.Errorf("failed, expected %v but got %v", tt.expectedOutput, output)
		}
	}
}
