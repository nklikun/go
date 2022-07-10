package main

// https://leetcode.cn/problems/binary-tree-preorder-traversal
func run144() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	nums := make([]int, 0)
	nums = append(nums, root.Val)
	nums = append(nums, preorderTraversal(root.Left)...)
	nums = append(nums, preorderTraversal(root.Right)...)
	return nums
}
