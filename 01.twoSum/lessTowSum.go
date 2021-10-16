package main
//Leetcode 1099：小于 K 的两数之和
import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{34,23,1,24,75,33,54,8}
	target := 60

	nums = []int{10, 20, 30}
	target = 15

	//start := time.Now()
	result := lessTwoSumViolence(nums, target)
	result = lessTwoSum(nums, target)
	//secs := time.Since(start).Seconds()
	//fmt.Printf("%.2fs  %7d", secs, result)
	fmt.Printf("%-7d\n", result)
}

func lessTwoSum(elems []int, target int) int {
	sort.Ints(elems)
	//fmt.Printf("%v\n", elems)
	result := -1
	for i, j := 0, len(elems) - 1; i < j; {
		sum := elems[i] + elems[j]
		if target > sum {
			i++
			result = sum
		} else {
			j--
		}
	}

	return result
}




func lessTwoSumViolence(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if sum := nums[i] + nums[j]; sum < target && result < sum {
				result = nums[i] + nums[j]
			}
		}
	}

	return  result
}