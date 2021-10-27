package _04_link

import (
	"fmt"
)

func main() {
	l1 := []int{2,4,3}
	l2 := []int{5,6,4}

	l1 = []int{9,9,9,9,9,9,9}
	l2 = []int{9,9,9,9}

	l1 = []int{2,4,9}
	l2 = []int{5,6,4,9}


	pl1 := makeLink(l1)
	pl2 := makeLink(l2)

	pl3 := addTwoNumbers(pl1, pl2)

	printLink(pl3)
}


type ListNode struct {
     Val int
     Next *ListNode
 }

func printLink(p *ListNode) {
	for; p != nil; {
		fmt.Printf("%d", p.Val)
		p = p.Next
	}

	fmt.Println()
}

 func makeLink(values []int) *ListNode {
 	var head, current *ListNode
	 head = makeNode(-1)
	 current = head

 	for _, elem := range values {
		current.Next = makeNode(elem)
		current = current.Next
	 }

	 return head.Next
 }

 func makeNode(value int) *ListNode {
	 t := new(ListNode) // equal to &tree{value: value}
	 t.Val = value
	 return t
 }
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, current *ListNode

	head = makeNode(-1)
	current = head
	carry := 0

	for; l1 != nil && l2 != nil; {
		sum := l1.Val + l2.Val + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}

		current.Next = makeNode(sum)
		current = current.Next

		l1 = l1.Next
		l2 = l2.Next
	}

	for ; l1 != nil ; {
		sum := l1.Val  + carry
		carry = 0
		if sum > 9 {
			carry = 1
			sum -= 10
		}

		current.Next = makeNode(sum)
		current = current.Next

		l1 = l1.Next


	}

	for ; l2 != nil ; {
		sum := l2.Val  + carry
		carry = 0

		if sum > 9 {
			carry = 1
			sum -= 10
		}

		current.Next = makeNode(sum)
		current = current.Next

		l2 = l2.Next
	}

	if carry == 1 {
		current.Next = makeNode(1)
		current = current.Next
	}

	return head.Next
}