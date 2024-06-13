package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}

	l := []*ListNode{}
	piv := head

	for piv.Next != nil {
		l = append(l, piv)
		piv = piv.Next
		for i := 0; i < len(l); i++ {
			if l[i] == piv {
				return l[i]
			}
		}
	}
	return nil
}

func main() {
	test := []ListNode{}

	t1 := ListNode{3, nil}
	t2 := ListNode{2, nil}
	t3 := ListNode{0, nil}
	t4 := ListNode{-4, nil}
	t1.Next = &t2
	t2.Next = &t3
	t3.Next = &t4
	t4.Next = &t2
	test = append(test, t1)

	for t := 0; t < len(test); t++ {
		fmt.Printf("%d", (*detectCycle(&test[t])).Val)
	}
}
