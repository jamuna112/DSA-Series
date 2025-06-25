package stringss

import "testing"

func TestIfPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"valid case", "racecar", true},
		{"Case-insensitive", "RaceCar", true},         
		{"Ignore punctuation and spaces", "A man, a plan, a canal: Panama", true}, 
		{"Case and punctuation", "No 'x' in Nixon", true},        
		{"Not a palindrome", "hello", false},                  
		{"Empty string is a palindrome", "", true},                         
		{"Digit + letter mismatch", "0P", false},                     
		{"classic palindrome", "Able was I, ere I saw Elba", true},
	}

	for _, tt := range tests {
		result := IsPalindrome(tt.input)

		if result != tt.expected {
			t.Errorf("isPalindrome(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
