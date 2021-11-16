package _04_link

//!+bench

import "testing"

// cache resolution
func findDuplicate_cacheResolution(nums []int) int {
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


// binary search resolution
func findDuplicate_binarySearchResolution(nums []int) int {
	len := len(nums)
	left, right := 1, len-1

	for ; left != right; {
		mid := left + (right - left) / 2
		cnt := 0
		for _, elem := range nums {
			if elem <= mid {
				cnt++
			}
		}

		if cnt > mid {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func findDuplicate(nums []int) int {
	r := 0

	for i := 0; i < 32; i++ {
		normal  := 0
		unusual := 0
		for j, elem := range nums {
			if j > 0 {
				normal += (j >> i) & 0x1
			}

			unusual += (elem >> i) & 0x1
		}

		if unusual > normal {
			r += 1 << i
		}

	}

	return r
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