package main

// https://leetcode.cn/problems/linked-list-cycle-ii
func run142() {
	l := buildNodeList([]int{1, 2, 3, 4, 5, 6})
	// 1-2-3-4-5-6-3-...
	l.Next.Next.Next.Next.Next.Next = l.Next.Next
	detectCycle(l)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil && slow != nil {
		if fast == slow {
			// slow比head慢了一个指针，所以slow先走一步
			slow = slow.Next
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return head
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return nil
}
