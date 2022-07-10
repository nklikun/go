package jianzhioffer

// https://leetcode.cn/problems/fan-zhuan-lian-biao-lcof/
func Run24() {}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	next := head.Next
	cur := head
	cur.Next = nil

	for next != nil {
		nextnext := next.Next
		next.Next = cur
		cur = next
		next = nextnext
	}
	return cur
}
