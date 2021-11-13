package _04_link

//!+bench


type CycleListNodeII struct {
	Val int
	Next *CycleListNodeII
}

func addHelperII(t *CycleListNodeII, value int) *CycleListNodeII {
	if t == nil {
		t = new(CycleListNodeII) // equal to &tree{value: value}
		t.Val = value
		return t
	}

	t.Next = addHelperII(t.Next, value)

	return t
}


func addII(elems []int, cyclePos int) *CycleListNodeII {
	links := make([]*CycleListNodeII, len(elems))

	for i, elem := range elems {
		t := new(CycleListNodeII) // equal to &tree{value: value}
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
func detectCycle(head *CycleListNodeII) *CycleListNodeII {
	if head == nil || head.Next == nil {
		return nil
	}

	pslow := head
	pfast := head
	for ;  pfast != nil; {

		pslow = pslow.Next
		if pfast.Next == nil {
			return nil
		}

		pfast = pfast.Next.Next

		if pfast == pslow {
			p := head
			for; p != pslow; {
				p = p.Next
				pslow = pslow.Next
			}
			return p
		}

	}

	return nil
}


//hash resolution
func detectCycleI(head *CycleListNodeII) *CycleListNodeII {
	cache := make(map[*CycleListNodeII]bool)

	p := head
	for; p != nil; p = p.Next {
		if _, ok := cache[p]; ok {
			return p
		}
		cache[head] = true
	}
	return nil
}


