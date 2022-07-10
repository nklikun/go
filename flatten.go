package main

// https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
func run114() {
	t := buildBinaryTreeByLevel([]int{1, 2, 5, 3, 4, null, 6})
	flatten(t)
	printBinaryTreeByLevel(t)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode) {
	var node *TreeNode
	var flattenHelper func(*TreeNode)
	flattenHelper = func(root *TreeNode) {
		if root == nil {
			return
		}
		flattenHelper(root.Right)
		flattenHelper(root.Left)
		root.Right = node
		root.Left = nil
		node = root
	}

	flattenHelper(root)
}
