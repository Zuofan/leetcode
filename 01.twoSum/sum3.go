package main

import (
	"fmt"
	"sort"
)

func main() {

	//sum3([]int{-1,0,1,2,-1,-4})
	sum3([]int{0,0,0,0})

	//[-1,0,1,2,-1,-4,-2,-3,3,0,4]
	//[[-4,0,4],[-4,1,3],[-3,-1,4],[-3,0,3],[-3,1,2],[-2,-1,3],[-2,0,2],[-1,-1,2],[-1,0,1]]

	//f :=[3][3]int {
	// {1,2,3},
	// {4,5,6},
	// {7,8,9},
	//}
	//
	//var m [][]int   //多维数组的切片，通过循环一下
	//for i, v :=range f{
	// fmt.Println(v[1:3])
	// m = append(m, f[i][1:3])
	// fmt.Println(m)
	//}
}


func sum3(elems []int) [][]int {
	sort.Ints(elems)

	var result [][]int

	cache := make(map[string]int)

	for i := 0; i < len(elems); i++ {
		if _, ok := cache[k(i, 0, 0)]; ok {
			continue
		}
		p, q, r := sum3Helper(elems, i)
		//fmt.Printf("%d %d %d result=%t\n",elems[i], elems[p], elems[q], r)
		if r {
			currentIndex := i
			if currentIndex > q {
				currentIndex, p, q = p, q, currentIndex
			} else if currentIndex > p {
				currentIndex, p, q = p, currentIndex, q
			}

			if _, ok := cache[k(elems[currentIndex], elems[p], elems[q])]; ok {
				continue
			}
			result = append(result, []int{elems[currentIndex], elems[p], elems[q]})
			cache[k(elems[currentIndex], elems[p], elems[q])] = currentIndex+p+q
			fmt.Printf("%d %d %d\n",currentIndex, p, q)
			cache[k(i, 0, 0)] = i
			cache[k(p, 0, 0)] = p
			cache[k(q, 0, 0)] = q
		}
	}
	fmt.Println(cache)
	return result
}

func k(i,p,q int) string {
	return fmt.Sprintf("%d%d%d", i, p, q)
}

func sum3Helper(elems []int, index int) (int, int, bool) {
	target := 0-elems[index]
	for i, j := 0, len(elems) - 1; i < j; {
		if i == index {
			i++
			continue
		} else if j == index {
			j--
			continue
		}
		sum := elems[i] + elems[j]
		if target > sum {
			i++
		} else if target < sum {
			j--
		} else {
			return i, j, true
		}
	}

	return 0, 0, false
}

