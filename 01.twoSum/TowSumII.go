package main

//Leetcode 167. Two Sum II - Input array is sorted
import (
	"fmt"
)

func main() {
	var result []int
	nums := []int{2, 11, 15, 7}
	target := 9

	nums = []int{2, 3, 4}
	target = 6
	target = 8

	result = twoSumII(nums, target)
	fmt.Printf("%-7d\n", result)
}

func twoSumII(nums []int, target int) []int {

	for i, j := 0, len(nums)-1; i < j; {
		sum := nums[i] + nums[j]
		if target > sum {
			i++
		} else if target < sum {
			j--
		} else {
			return []int{i + 1, j + 1}
		}
	}

	return nil
}
