package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	mid := 0
	piv := head
	for piv != nil {
		mid += 1
		piv = piv.Next
	}
	for i := 0; i < mid/2; i++ {
		head = head.Next
	}
	return head
}

func main() {

}
