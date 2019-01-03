// https://leetcode.com/problems/add-two-numbers/
package main

// ListNode is Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	head := &ListNode{Val: 0}
	pointer := head
	next := 0

	for {
		pointer.Val = l1.Val + l2.Val + next
		next = pointer.Val / 10
		pointer.Val %= 10

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			break
		}
		if l1 == nil {
			l1 = &ListNode{Val: 0}
		}
		if l2 == nil {
			l2 = &ListNode{Val: 0}
		}

		pointer.Next = &ListNode{Val: 0}
		pointer = pointer.Next
	}

	if next != 0 {
		pointer.Next = &ListNode{Val: next}
	}

	return head
}
