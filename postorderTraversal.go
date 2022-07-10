package main

// https://leetcode.cn/problems/binary-tree-postorder-traversal
func run145() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := make([]int, 0)
	nums = append(nums, postorderTraversal(root.Left)...)
	nums = append(nums, postorderTraversal(root.Right)...)
	nums = append(nums, root.Val)
	return nums
}
