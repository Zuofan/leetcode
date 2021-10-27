package _04_link

//!+bench

import (
	"strconv"
	"testing"
)


func AddBinary(a string, b string) string {
	r := ""
	var i,j, carry int

	carry = 0

	for i, j = len(a) - 1, len(b) -1 ; i >= 0 && j >= 0; {
		i1, _ := strconv.Atoi(string(a[i]))
		j1, _ := strconv.Atoi(string(b[j]))
		sum := i1 + j1 + carry
		carry = 0
		if sum > 1 {
			carry = 1
			sum %= 2
		}

		r = strconv.Itoa(sum) + r

		i--
		j--
	}

	for ; i >= 0;  {
		i1, _ := strconv.Atoi(string(a[i]))
		sum := i1  + carry
		carry = 0
		if sum > 1 {
			carry = 1
			sum %= 2
		}

		r = strconv.Itoa(sum) + r
		i--
	}

	for ; j >= 0; {
		j1, _ := strconv.Atoi(string(b[j]))
		sum := j1  + carry
		carry = 0
		if sum > 1 {
			carry = 1
			sum %= 2
		}

		r = strconv.Itoa(sum) + r
		j--
	}

	if carry != 0 {
		r = strconv.Itoa(1) + r
	}
	return r
}


//!+test
func TestAddBinary(t *testing.T) {
	var tests = []struct {
		a, b string
		want  string
	}{
		{"111", "1", "1000"},
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, test := range tests {
		if got := AddBinary(test.a, test.b); got != test.want {
			t.Errorf("AddBinary(%q, %q) = %q, got=%q\n", test.a, test.b, test.want, got)
		}
	}
}
//!-test