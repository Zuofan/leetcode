package main
//Leetcode ：Subarray Sum Equals K
import (
	"log"
	"os"
)

func main() {
	//violence
	assertEqual(2, subarraySumViolence([]int{1,2,3}, 3))
	assertEqual(4, subarraySumViolence([]int{1,2,1,2,1}, 3))
	assertEqual(2, subarraySumViolence([]int{1,1,1}, 2))
	assertEqual(1, subarraySumViolence([]int{28,54,7,-70,22,65,-6}, 100))

}

func assertEqual(expect, actual  int) bool {
	if expect == actual  {
		return true
	}

	log.Fatalf("expect %d real:%d\n", expect, actual )
	os.Exit(1)
	return false
}

/*func assertEqual(t *testing.T, expect, actual interface{}) {
	t.Helper()
	if a != b {
		t.Errorf("Not Equal. %d %d", a, b)
	}
}*/

func subarraySumViolence(nums []int, k int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		currentSum := nums[i]
		if currentSum == k {
			result++
		}
		for j := i+1; j < len(nums); j++ {
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