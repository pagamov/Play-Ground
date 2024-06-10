package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func lenListNode(l *ListNode) int {
	var count int = 0
	var piv *ListNode = l
	for piv != nil {
		fmt.Print(piv.Val, " ")
		count += 1
		piv = piv.Next
		// fmt.Println(piv)
	}
	return count
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if lenListNode(l1) >= lenListNode(l2) {
		var carry int = 0
		var piv1 *ListNode = l1
		var piv2 *ListNode = l2

		for true {
			piv1.Val += piv2.Val + carry
			if carry == 1 {
				carry = 0
			}
			if piv1.Val >= 10 {
				carry = 1
				piv1.Val -= 10
			}
			if piv1.Next == nil {
				if carry == 1 {
					piv1.Next = &ListNode{Val: 1, Next: nil}
				}
				return l1
			} else {
				if piv1.Next == nil {
					piv1.Next = &ListNode{Val: 0, Next: nil}
				}
				piv1 = piv1.Next
				if piv2.Next == nil {
					piv2.Next = &ListNode{Val: 0, Next: nil}
				}
				piv2 = piv2.Next
			}
		}
	} else if lenListNode(l1) < lenListNode(l2) {
		var carry int = 0
		var piv1 *ListNode = l2
		var piv2 *ListNode = l1

		for true {
			piv1.Val += piv2.Val + carry
			if carry == 1 {
				carry = 0
			}
			if piv1.Val >= 10 {
				carry = 1
				piv1.Val -= 10
			}

			if piv1.Next == nil {
				if carry == 1 {
					piv1.Next = &ListNode{Val: 1, Next: nil}
				}
				return l2
			} else {
				if piv1.Next == nil {
					piv1.Next = &ListNode{Val: 0, Next: nil}
				}
				piv1 = piv1.Next
				if piv2.Next == nil {
					piv2.Next = &ListNode{Val: 0, Next: nil}
				}
				piv2 = piv2.Next
			}
		}
	}
	return l1
}

func main() {
	// l1 = [2,4,3], l2 = [5,6,4]
	var l1 ListNode = ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}}}
	var l2 ListNode = ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}}}

	var res *ListNode = addTwoNumbers(&l1, &l2)

	fmt.Println(res)
}
