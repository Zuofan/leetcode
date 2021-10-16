package main
//Leetcode ：两数之和
import (
	"fmt"
)

func main() {

	elems := []int{2,8,7,15,19}
	sum := 17
	result := twoSumStand(elems, sum)
	fmt.Printf("%q\n", result)
/*	ages := map[string]int{
		"alice":31,
		"charlie":34,
	}

	for i, j := range ages {
		fmt.Printf("%q %d\n", i, j)
	}

	i, name := ages["name"]
	fmt.Printf("%q %t\n", i, name)*/
}


func twoSumMy(elems []int, sum int) {
	helper := make(map[int]int)
	helper[elems[0]] = 0 //special because of default value is zero

	for i, elem := range elems[1:] {
		if _, ok := helper[sum-elem]; ok {
			fmt.Printf("[%d, %d]\n", helper[sum-elem], i+1)
		}
	}

}

func twoSumStand(elems []int, sum int) []int {
	cache := make(map[int]int)
	for i := 0; i < len(elems); i++ {
		left := sum - elems[i]
		if _, ok := cache[left]; ok {
			return []int{cache[left], i}
		}

		cache[elems[i]] = i
	}

	return nil
}

func twoSumNormal(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}