package main

import (
	"fmt"
	"sort"
)

func main() {

	//sum3([]int{-1,0,1,2,-1,-4})
	fmt.Println(threeSum([]int{0,0,0,0}))

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


func threeSum(elems []int) [][]int {
	sort.Ints(elems)
	var results [][]int

	for i:= 0; i < len(elems); i++ {
		if i > 0 && elems[i] == elems[i-1] {
			continue
		}

		for j:= i+1; j < len(elems); j++ {

			if j > i+1 && elems[j] == elems[j-1] {
				continue
			}

			for m := j+1; m < len(elems); m++ {
				if m > j+1 && elems[m] == elems[m-1] {
					continue
				}

				if elems[i] + elems[j] + elems[m] == 0 {
					results = append(results, []int{elems[i], elems[j], elems[m]})
					continue
				}

			}
		}
	}

	return results
}