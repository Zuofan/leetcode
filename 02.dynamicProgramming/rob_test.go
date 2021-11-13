package _02_dynamicprogramming

//func main() {
//	fmt.Println(rob([]int{1,2,3,1}))
//}

func rob(costs []int) int {
	if len(costs) == 1 {
		return costs[0]
	}

	p, q, r := costs[0], costs[1], 0
	if p > q {
		r = p
	} else {
		r = q
	}

	if len(costs) == 2 {
		return r
	}

	q = r

	for i := 2; i < len(costs); i++ {

		r = p + costs[i]
		if r < q {
			r = q
		}

		p = q
		q = r

	}

	return r
}

