package jianzhioffer

// https://leetcode.cn/problems/dui-cheng-de-er-cha-shu-lcof/
func Run28() {}

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

	return isSymmetric_helper(root.Left, root.Right)
}

func isSymmetric_helper(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil {
		return false
	}

	return root1.Val == root2.Val &&
		isSymmetric_helper(root1.Left, root2.Right) &&
		isSymmetric_helper(root1.Right, root2.Left)
}
