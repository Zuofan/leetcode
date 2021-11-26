package main

import "fmt"

func main() {
	var name string
	fmt.Printf("请问你是谁? ")
	fmt.Scanf("%s", &name)

	fmt.Printf("你好, %s, 欢迎你使用Go!", name)
}