package _02_dynamicprogramming

//!+bench

import "testing"

func maxSubArray(nums []int) int {
	len := len(nums)
	cache := make([]int, len)

	max := nums[0]
	cache[0] = nums[0]
	for i := 1; i < len; i++ {
		temp := cache[i-1]
		if cache[i-1] + nums[i] > temp {
			temp = cache[i-1] + nums[i]
		}
		cache[i] = temp

		if max < temp {
			max = temp
		}
	}

	return max
}


func maxSubArrayV1(nums []int) int {
	len := len(nums)
	cache := make([]int, len)

	s := 0
	for i := 0; i < len; i++ {
		s += nums[i]
		cache[i] = s
	}

	max := 0
	for i := 0; i < len; i++ {
		for j:= i; j < len; j++ {
			if max < (cache[j] - cache[i] + nums[i]) {
				max = cache[j] - cache[i]+ nums[i]
			}
		}
	}


	return max
}

//!+test
func TestMaxSubArray(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{-2,1,-3,4,-1,2,1,-5,4}, 6},
		{[]int{1}, 1},
		{[]int{5,4,-1,7,8}, 23},
		{[]int{-1}, -1},
	}
	for _, test := range tests {
		if got := maxSubArray(test.input); got != test.want {
			t.Errorf("xxx(%q) = %v want=%v", test.input, got, test.want)
		}
	}
}
//!-test