package stringss

import (
	"reflect"
	"testing"
)

func TestRemoveOccurances(t *testing.T) {

	tests := []struct {
		input          string
		part           string
		expectedOutput string
	}{
		{"daabcbaabcbc", "abc", "dab"},
		{"axxxxyyyyb", "xy", "ab"},
	}

	for _, tt := range tests {
		output := RemoveOccurrences(tt.input, tt.part)

		if !reflect.DeepEqual(output, tt.expectedOutput) {
			t.Errorf("Failed. Expected %v but got %v", tt.expectedOutput, output)
		}
	}
}
