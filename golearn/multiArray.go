package main

import (
	"fmt"
)

func main() {

	// 首先make第一维的大小
	var arr = make([][]int, 2)
	// 然后分别对其中的进行make
	for i := range arr{
		arr[i] = make([]int, 3)
	}

	fmt.Println(arr)


	mapResults := make(map[int]string)
	var arrResults [][]string

	count := 5
	for i := 0; i < count; i++ {
		valueStr := fmt.Sprintf("this is %d", i)
		mapResults[i] = valueStr
		var tmpArr []string

		for j := 0; j < 15; j++ {
			tmpArr = append(tmpArr, "a")

		}

		arrResults = append(arrResults, tmpArr)

	}

	fmt.Println(mapResults)
	fmt.Println(arrResults)
}