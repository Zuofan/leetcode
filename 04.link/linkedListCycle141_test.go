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
	var r *CycleListNode
	var cyclePoint *CycleListNode
	var tail *CycleListNode

	for i, elem := range elems {
		r = addHelper(r, elem)
		if i == cyclePos {
			cyclePoint = r
		}

		if len(elems) - 1 == i {
			tail = r
		}
	}

	tail.Next = cyclePoint

	return r
}

func hasCycle(head *CycleListNode) bool {
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
func TestHasCycle(t *testing.T) {
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