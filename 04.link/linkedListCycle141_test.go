package _04_link

//!+bench

import "testing"


type CycleListNode struct {
	Val int
	Next *CycleListNode
}

func addHelper(t *CycleListNode, value int) *CycleListNode {
	if t == nil {
		t = new(CycleListNode) // equal to &tree{value: value}
		t.Val = value
		return t
	}

	t.Next = addHelper(t.Next, value)

	return t
}


func add(elems []int, cyclePos int) *CycleListNode {
	links := make([]*CycleListNode, len(elems))

	for i, elem := range elems {
		t := new(CycleListNode) // equal to &tree{value: value}
		t.Val = elem

		links[i] = t
	}


	for i, _ := range elems {
		if i < len(elems) - 1 {
			links[i].Next = links[i+1]
		} else if cyclePos >= 0 {
			links[len(elems) - 1].Next = links[cyclePos]
		}

	}

	return links[0]
}


//Floyd Cycle Detection Algorithm
func hasCycle(head *CycleListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	pslow := head
	pfast := head.Next.Next
	for ; pslow != pfast; {
		if pslow.Next == nil || pfast == nil {
			break
		}

		pslow = pslow.Next

		if pfast != nil && pfast.Next != nil && pfast.Next.Next != nil {
			pfast = pfast.Next.Next
		}
	}

	return pslow == pfast
}

func TestHasCycle(t *testing.T) {
	var tests = []struct {
		input *CycleListNode
		want  bool
	}{
		{add([]int{3,2,0,-4}, 1), true},
		{add([]int{1,2}, 0), true},
		{add([]int{1}, -1), false},
		{add([]int{1, 2}, -1), false},
	}
	for j, test := range tests {
		if got := hasCycle(test.input); got != test.want {
			t.Errorf("%d:hasCycle(%v) = %v", j, test.input, got)
		}
	}
}

//hash resolution
func hasCycleV1(head *CycleListNode) bool {
	cache := make(map[*CycleListNode]bool)

	for; head != nil; head = head.Next {
		if _, ok := cache[head]; ok {
			return true
		}
		cache[head] = true
	}
	return false
}


//!+test
func TestHasCycleV1(t *testing.T) {
	var tests = []struct {
		input *CycleListNode
		want  bool
	}{
		{add([]int{3,2,0,-4}, 1), true},
		{add([]int{1,2}, 0), true},
		{add([]int{1}, -1), false},
	}
	for _, test := range tests {
		if got := hasCycle(test.input); got != test.want {
			t.Errorf("hasCycle(%v) = %v", test.input, got)
		}
	}
}
//!-test