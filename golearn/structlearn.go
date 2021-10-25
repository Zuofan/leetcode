package main

import (
	"fmt"
)



type tree struct {
	value int
	left, right *tree
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree) // equal to &tree{value: value}
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}

func Sort(values []int)  {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}

	appendValues(values[:0], root)

}

func appendValues(values []int , t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

func main() {
	values := []int{2, 5, 3, 1, 7}
	Sort(values)
	fmt.Println(values)
}