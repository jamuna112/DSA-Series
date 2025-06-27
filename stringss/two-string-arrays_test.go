package stringss

import "testing"

func TestTwoArrayStrings(t *testing.T) {

	tests := []struct {
		input1         []string
		input2         []string
		expectedOutput bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"ar", "bc"}, []string{"ab", "b"}, false},
		{[]string{"abcddefg"}, []string{"abc", "d", "defg"}, true},
		{[]string{"hello", ""}, []string{"he", "llo"}, true},
		{[]string{}, []string{}, true},
	}

	for _, tt := range tests {
		output := ArrayStringsAreEqual(tt.input1, tt.input2)

		if output != tt.expectedOutput {
			t.Errorf("failed, expected %v but got %v", tt.expectedOutput, output)
		}

	}

}
