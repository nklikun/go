package main

import "fmt"

// https://leetcode.cn/problems/sum-of-left-leaves
func run404() {
	t := buildBinaryTreeByLevel([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(sumOfLeftLeaves(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		l = root.Left.Val
	} else {
		l = sumOfLeftLeaves(root.Left)
	}

	return l + sumOfLeftLeaves(root.Right)
}
