package _04_link

import (
	"testing"
)


func singleNumberII(nums []int) int {
	cache := make(map[int]int)


	for _, elem := range nums {
		cache[elem]++
	}

	for elem, occ := range cache {
		if occ == 1 {
			return elem
		}
	}

	return 0
}

//!+test
func TestSingleNumberII(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2,2,1, 2}, 1},
		{[]int{4,1,2,1,2,1,2}, 4},
	}
	for _, test := range tests {
		if got := singleNumberII(test.input); got != test.want {
			t.Errorf("singleNumber(%q) = %v", test.input, got)
		}
	}
}
//!-test