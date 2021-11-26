package main

import "fmt"

func main() {
	var first int
	var second int

	fmt.Printf("First Number ")
	fmt.Scanf("%d", &first)

	fmt.Printf("Second Number ")
	fmt.Scanf("%d", &second)

	fmt.Printf("%d + %d = %d\n", first, second, first + second)
	fmt.Printf("%d - %d = %d\n", first, second, first - second)
	fmt.Printf("%d * %d = %d\n", first, second, first * second)
	fmt.Printf("%d / %d = %d\n", first, second, first / second)
}