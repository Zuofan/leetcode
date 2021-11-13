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
		//if got := xxxx(test.input); got != test.want {
		//	t.Errorf("xxx(%q) = %v want=%v", test.input, got, test.want)
		//}
	}
}
//!-test