package main

import (
"fmt"
)

func main() {
	s := "abcde"
	b := []byte(s)
	sum := 0
	for _, ch := range b {
		sum += int(ch)
	}

	fmt.Printf("%d", sum)
}
