package stringss

import "testing"

func TestRlEncoding(t *testing.T) {
	tests := []struct {
		name        string
		inputStr    string
		expectedStr string
	}{
		{"bunch of characters", "aaabbccccd", "a3b2c4d1"},
		{"empty string", "", ""},
		{"single character", "a", "a1"},
		{"similar characters", "aaaaaa", "a6"},
		{"input numbers", "111222333", "132333"},
		{"non-alphabetical characters", "!!!###$$", "!3#3$2"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := RlEncoding(tt.inputStr)

			if output != tt.expectedStr {
				t.Errorf("%s test failed, expected %v but got %v", tt.name, tt.expectedStr, output)
			}
		})
	}
}
