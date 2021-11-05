package _04_link

import (
	"sort"
	"testing"
)

//bit operator:great
func missingNumber(nums []int) int {
	len := len(nums)
	n := 0

	for i := 0; i < len ; i++ {
		n ^= nums[i]
		n ^= i
	}

	n ^= len

	return n
}


//math method:great!!!
func missingNumberV3(nums []int) int {
	len := len(nums)
	sum := 0
	for i := 0; i < len ; i++ {
		sum += nums[i]
	}

	return len * (len+1) / 2 - sum
}


//hash resolution
func missingNumberV2(nums []int) int {
	cache := make(map[int]int)
	len := len(nums)

	for i := 0; i < len ; i++ {
		cache[nums[i]] = i
	}

	for i:= 0; i <= len; i++ {
		if _, ok := cache[i]; !ok {
			return i
		}
	}

	return 0
}

//sort method:optimization
func missingNumberV11(nums []int) int {
	sort.Ints(nums)
	len := len(nums)

	for i := 0; i < len ; i++ {
		if nums[i] != i {
			return i
		}
	}

	return len
}

//sort method:normal
func missingNumberV1(nums []int) int {
	sort.Ints(nums)
	nextElem := 0
	for i := 0; i < len(nums) ; i++ {
		if nums[i] != nextElem {
			return nextElem
		}
		nextElem += 1
	}

	return nextElem
}

//!+test
func TestMissingNumber(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{3,0,1}, 2},
		{[]int{0,1}, 2},
		{[]int{9,6,4,2,3,5,7,0,1}, 8},
		{[]int{0}, 1},
	}
	for _, test := range tests {
		if got := missingNumber(test.input); got != test.want {
			t.Errorf("missingNumber(%q) want=%v got=%v", test.input, test.want, got)
		}
	}
}
//!-test