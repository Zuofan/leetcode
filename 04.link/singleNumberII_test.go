package _04_link

import (
	"testing"
)


func singleNumberII(nums []int) int {
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

	r := 0
	for i := 0; i < 32; i++ {
		sum := 0
		for _, elem := range nums {
			sum += (elem >> i) & 0x1
			//fmt.Printf("elem=%032b, sum=%b\n", elem, sum)
		}
		if sum % 3 == 1 {
			if i == 31 {
				r -= 1 << i
			} else {
				r |= 1 << i
			}
		}
	}
	return r
}
func singleNumberIIV2(nums []int) int {
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
<<<<<<< HEAD
		{[]int{4,1,2,1,2, 2,1}, 4},
		{[]int{-2,-2,1,1,4,1,4,4,-4,-2}, -4},
=======
		{[]int{4,1,2,1,2,1,2}, 4},
>>>>>>> c5350a35f2760ddaaea1181febf8911cb9d0c61c
	}
	for _, test := range tests {
		if got := singleNumberII(test.input); got != test.want {
			t.Errorf("singleNumber(%q) = %v", test.input, got)
		}
	}
}
//!-test