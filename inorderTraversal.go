package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-inorder-traversal
func run94() {
	t := buildBinaryTreeByLevel([]int{1, NULL, 2, NULL, 3, 4})
	n := inorderTraversal(t)
	fmt.Println(n)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nums := make([]int, 0)

	nums = append(nums, inorderTraversal(root.Left)...)
	nums = append(nums, root.Val)
	nums = append(nums, inorderTraversal(root.Right)...)

	return nums
}
