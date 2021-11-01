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
	r := 0
	for i := 0; i < 32; i++ {
		sum := 0
		for _, elem := range nums {
			sum += (elem >> i) & 0x1
		}

		r += (sum % 2) << i
	}

	return r

	//r := int32(0)
	//for i := 0; i < 32; i++ {
	//	sum := int32(0)
	//	for _, elem := range nums {
	//		sum += (int32(elem) >> i) & 0x1
	//	}
	//
	//	r |= (sum % 3) << i
	//}
	//
	//return int(r)
}

func singleNumberVV1(nums []int) int {
	cache := make(map[int]int)

	for _, elem := range nums {
		cache[elem] += 1
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
	//
	//return r
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