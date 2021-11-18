package _04_link



//!+bench

import "testing"


func findTheDifference(s string, t string) byte {
	sum := 0
	for _, ch := range s {
		sum += int(ch)
	}

	for _, ch := range t {
		sum -= int(ch)
	}

	if sum < 0 {
		sum = 0 - sum
	}

	return byte(sum)
}


//!+test
func TestFindTheDifference(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		want  byte
	}{
		{"abcd", "abecd", 'e'},
		{"", "y", 'y'},
		{"a", "aa", 'a'},
		{"ae", "aea", 'a'},
	}
	for _, test := range tests {
		if got := findTheDifference(test.input1, test.input2); got != test.want {
			t.Errorf("xxx(%q, %q) = %v want=%v", test.input1, test.input2, got, test.want)
		}
	}
}
//!-test
