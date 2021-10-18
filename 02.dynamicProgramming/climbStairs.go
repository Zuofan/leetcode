package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	fibs := []int{0, 1}
	fibResult := 0

	for i:= 2; i < n+1; i++ {
		fibResult = fibs[i-2] + fibs[i-1]
		fibs = append(fibs, fibResult)
	}

	return fibResult
}

