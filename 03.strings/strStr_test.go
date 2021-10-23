package _3_strings

import "testing"

//!+test
func TestStringsStrStr(t *testing.T) {
	var tests = []struct {
		input1, input2 string
		want  int
	}{
		{"", "", 0},
		{"hello", "ll",2},
	}
	for _, test := range tests {
		if got := strStr(test.input1, test.input2); got != test.want {
			t.Errorf("strStr(input1=%q input2=%q) got = %d want = %d\n", test.input1, test.input2, got, test.want)
		}
	}
}
//!-test


func strStr(haystack string, needle string) int {
	r := -1
	n1, n2 := len(haystack), len(needle)

	if n2 == 0 {
		r = 0
	}

	for i:=0; i < n1; i++ {
		if n2+i > n1 {
			break
		}

		if haystack[i:n2+i] == needle {
			r = i
			break
		}
	}

	return r
}
