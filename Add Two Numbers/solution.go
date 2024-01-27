// https://leetcode.com/problems/add-two-numbers

package solution

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyList := &ListNode{Val: 0}
	it := dummyList
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}

		sum := carry + x + y
		carry = sum / 10 // save carry if the addition of x + y is greater than 10

		it.Next = &ListNode{Val: sum % 10}
		it = it.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	return dummyList.Next
}
