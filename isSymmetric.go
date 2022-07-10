package main

// https://leetcode.cn/problems/symmetric-tree
func run101() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return compareTreeLeftToRight(root.Left, root.Right)
}

func compareTreeLeftToRight(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil || left.Val != right.Val {
		return false
	}

	return compareTreeLeftToRight(left.Left, right.Right) && compareTreeLeftToRight(left.Right, right.Left)
}
