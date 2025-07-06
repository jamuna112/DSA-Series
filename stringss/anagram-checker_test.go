package stringss

import "testing"

func TestAnagramChecker(t *testing.T) {

	tests := []struct {
		inputStr1 string
		inputStr2 string
		expected  bool
	}{
		{"read", "dear", true},
		{"listen", "silent", true},
		{"david", "davinder", false},
		{"app", "ap", false},
		{"", "", true},
		{"e", "e", true},
	}

	for _, tt := range tests {
		t.Run(tt.inputStr1+" - "+tt.inputStr2, func(t *testing.T) {
			output := AnagramChecker(tt.inputStr1, tt.inputStr2)

			if output != tt.expected {
				t.Errorf("test failed, expected %v but got %v", tt.expected, output)
			}
		})
	}
}
