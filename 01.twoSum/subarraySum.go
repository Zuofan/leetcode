package main

/*
560. Subarray Sum Equals K

Given an array of integers nums and an integer k, return the total number of continuous subarrays whose sum equals to k.

Example 1:

Input: nums = [1,1,1], k = 2
Output: 2
Example 2:

Input: nums = [1,2,3], k = 3
Output: 2


Constraints:

1 <= nums.length <= 2 * 104
-1000 <= nums[i] <= 1000
-107 <= k <= 107


*/
import (
	"fmt"
	"log"
	"os"
)

func main() {
	//efficient
	//assertEqual(2, subarraySum([]int{1,2,3}, 3))
	//assertEqual(4, subarraySum([]int{1,2,1,2,1}, 3))
	assertEqual(2, subarraySum([]int{1,1,1}, 2))
	//assertEqual(1, subarraySum([]int{28,54,7,-70,22,65,-6}, 100))
	//assertEqual(0, subarraySum([]int{1}, 0))
	//assertEqual(1, subarraySum([]int{-1,-1,1}, 0))
	//assertEqual(3, subarraySum([]int{1,-1,0}, 0))

	////brute force
	//assertEqual(2, subarraySumBruteForce([]int{1,2,3}, 3))
	//assertEqual(4, subarraySumBruteForce([]int{1,2,1,2,1}, 3))
	//assertEqual(2, subarraySumBruteForce([]int{1,1,1}, 2))
	//assertEqual(1, subarraySumBruteForce([]int{28,54,7,-70,22,65,-6}, 100))

}

func subarraySum(nums []int, k int) int {
	//ver1:wrong!
	// result := 0
	// length := len(nums)

	// for i := 0; i < length; i++ {
	// 	currentSum := nums[i]
	// 	if currentSum == k {
	// 		result++
	// 	}
	// 	for j := i+1; j < length; j++ {
	// 		if currentSum + nums[j] == k  {
	// 			currentSum += nums[j]
	// 			result++
	// 		} else {
	// 			currentSum += nums[j]
	// 		}
	// 	}
	// }

	// return result

	//ver2:wrong!
	// result := 0
	// length := len(nums)
	// prefixSum := make(map[int]int)
	// prefixSum[0] = 0
	// currentSum := 0

	// existSum := make(map[int]int)
	// existSum[0] = 0

	// for j := 0; j < length; j++ {
	// 	currentSum += nums[j]
	// 	prefixSum[j+1] = currentSum
	// 	existSum[currentSum] = j+1

	// 	//fmt.Printf("prefixSum[%d]=%d\n", j+1, prefixSum[j+1])
	// 	if i, ok := existSum[prefixSum[j+1] - k]; ok && i < j+1 {
	// 		result++
	// 	}
	// }

	// return result

	//ver3:wrong!
	//count := 0
	//length := len(nums)
	//
	//existSum := make(map[int]int)
	//existSum[0] = 0
	//prefixSum := 0
	//
	//for j := 0; j < length; j++ {
	//	prefixSum += nums[j]
	//	if _, ok := existSum[prefixSum - k]; ok {
	//		count += 1
	//	}
	//
	//	existSum[prefixSum] = 0
	//}
	//
	//return count

	//ver4:right todo:why
	count := 0
	length := len(nums)

	sums := make(map[int]int)
	sums[0] = 1//?
	prefixSum := 0

	for j := 0; j < length; j++ {
		prefixSum += nums[j]
		count += sums[prefixSum - k]
		fmt.Printf("prefixSum=%d sums[%d]=%d, count=%d\n", prefixSum, prefixSum, sums[prefixSum],count)
		sums[prefixSum] = sums[prefixSum] + 1
	}

	return count
}

func subarraySumBruteForce(nums []int, k int) int {
	result := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		currentSum := nums[i]
		if currentSum == k {
			result++
		}
		for j := i+1; j < length; j++ {
			if currentSum + nums[j] == k  {
				currentSum += nums[j]
				result++
			} else {
				currentSum += nums[j]
			}
		}
	}

	return result
}

func assertEqual(expect, actual  int) bool {
	if expect == actual  {
		return true
	}

	log.Fatalf("expect %d real:%d\n", expect, actual )
	os.Exit(1)
	return false
}