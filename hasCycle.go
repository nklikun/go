package main

// https://leetcode.cn/problems/linked-list-cycle
func run141() {
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	next1, next2 := head, head.Next
	for next1 != nil && next2 != nil && next2.Next != nil {
		if next1 == next2 {
			return true
		}
		next1 = next1.Next
		next2 = next2.Next.Next
	}
	return false
}
