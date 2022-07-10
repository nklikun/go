package main

// https://leetcode.cn/problems/invert-binary-tree
func run226() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := invertTree(root.Left)
	root.Left = invertTree(root.Right)
	root.Right = tmp

	return root
}
