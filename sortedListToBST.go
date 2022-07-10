package main

// https://leetcode.cn/problems/convert-sorted-list-to-binary-search-tree
func run109() {
	head := buildNodeList([]int{1, 2})
	t := sortedListToBST(head)
	printBinaryTreeByLevel(t)
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	mid := head
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		mid = mid.Next
		fast = fast.Next.Next
	}
	prev := mid
	mid = mid.Next
	node := &TreeNode{Val: mid.Val}
	right := mid.Next
	prev.Next = nil
	node.Left = sortedListToBST(head)
	node.Right = sortedListToBST(right)

	return node
}
