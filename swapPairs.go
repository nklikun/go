package main

// https://leetcode.cn/problems/swap-nodes-in-pairs
func run24() {
	list := buildNodeList([]int{1})
	ret := swapPairs(list)

	printNodeList(ret)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	node1 := &ListNode{Val: 0, Next: head} // eg. 0-1-2-3-4, and use 0 to change 1and2
	node1Head := node1

	if head == nil || head.Next == nil {
		return head
	}

	for node1.Next != nil && node1.Next.Next != nil {
		node2 := node1.Next          // node1: 0-1-2-3-4, node2: 1-2-3-4
		node1.Next = node2.Next      // node1: 0-2-3-4, node2: 1-2-3-4
		node2.Next = node2.Next.Next // node1: 0-2-3-4, node2: 1-3-4
		node1.Next.Next = node2      // node1: 0-2-1-3-4, node2: 1-3-4
		node1 = node2                // node1: 1-3-4, nodeHead: 0-2-1-3-4
	}

	return node1Head.Next
}
