package main

// https://leetcode.cn/problems/insert-into-a-binary-search-tree
func run701() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	new := &TreeNode{Val: val}
	if root == nil {
		return new
	}
	node := root
	for node != nil {
		if val < node.Val {
			if node.Left == nil {
				node.Left = new
				break
			}
			node = node.Left
		} else {
			if node.Right == nil {
				node.Right = new
				break
			}
			node = node.Right
		}
	}

	return root
}
