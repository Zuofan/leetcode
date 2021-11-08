package _04_link

//!+bench

import "testing"


func findDuplicate(nums []int) int {
	cache := make(map[int]int)

	for _, elem := range nums {
		cache[elem]++
	}

	for _, elem := range nums {
		if count, _ := cache[elem]; count > 1 {
			return elem
		}
	}

	return 0
}


//!+test
func TestFindDuplicate(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{1,3,4,2,2}, 2},
		{[]int{3,1,3,4,2}, 3},
		{[]int{1,1}, 1},
		{[]int{1,1,2}, 1},
		{[]int{2,2,2,2,2}, 2},
	}
	for _, test := range tests {
		if got := findDuplicate(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}
//!-test