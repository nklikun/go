package main

// https://leetcode.cn/problems/remove-nth-node-from-end-of-list
func run19() {
	head := buildNodeList([]int{1, 2})
	removeNthFromEnd(head, 2)
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	nodeN := head
	for node := head; nil != node; node = node.Next {
		i++
		// n+1, we need find out the node before n node, and make next pointer to the node next to n
		if i > n+1 {
			nodeN = nodeN.Next
		}
	}

	if i < 1 {
		return nil
	} else if i < n {
		return head
	} else if i == n {
		return head.Next
	}

	nodeN.Next = nodeN.Next.Next
	return head
}
