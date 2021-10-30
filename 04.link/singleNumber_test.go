package _04_link

import (
	"sort"
	"testing"
)

func singleNumberV2(nums []int) int {
	r := 0

	for _, elem := range nums {
		r ^= elem
	}

	return r
}

func singleNumber(nums []int) int {
	cache := make(map[int]int)

	for _, elem := range nums {
		if _, ok := cache[elem]; ok {
			cache[elem] += 1
		} else {
			cache[elem] = 1
		}
	}

	for _, elem := range nums {
		if cache[elem] == 1 {
			return elem
		}
	}

	return 0
}

func singleNumberV1(nums []int) int {
	sort.Ints(nums)
	i := 0
	for ; i + 1 < len(nums);  {
		if nums[i] == nums[i+1] {
			i += 2
		} else {
			break
		}
	}

	return nums[i]
}


//!+test
func TestSingleNumber(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{2,2,1}, 1},
		{[]int{4,1,2,1,2}, 4},
	}
	for _, test := range tests {
		if got := singleNumber(test.input); got != test.want {
			t.Errorf("singleNumber(%q) = %v", test.input, got)
		}
	}
}
//!-test