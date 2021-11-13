package _02_dynamicprogramming

//func main() {
//	fmt.Println(climbStairs(5))
//}
func climbStairs(n int) int {
	if n == 1 || n== 2 {
		return n
	}

	p,q,r := 1, 2, 0

	for i:= 3; i < n+1; i++ {

		r = p + q
		p = q
		q = r
	}

	return r
}

func climbStairsV1(n int) int {
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

