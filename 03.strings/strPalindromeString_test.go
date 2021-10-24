package _3_strings

//!+bench

import (
	"strings"
	"testing"
)


func palindrome(str, substr string) bool {
	s := str + str[:len(str)-1]

	return strings.Index(s, substr) != -1
}


//!+test
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		str, substr string
		want  bool
	}{
		{"AABBCD", "CDAA", true},
		{"AABBCD", "CDAB", false},
	}
	for _, test := range tests {
		if got := palindrome(test.str, test.substr); got != test.want {
			t.Errorf("palindrome(%q, %q) = %v got = %v\n", test.str, test.substr, test.want, got)
		}
	}
}
//!-test