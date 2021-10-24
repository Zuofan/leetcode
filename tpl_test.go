package leetcode

//!+bench

import "testing"


//func target(str string) bool {
//	return true
//}


//!+test
func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
	}
	for _, test := range tests {
		//if got := target(test.input); got != test.want {
		//	t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		//}
	}
}
//!-test