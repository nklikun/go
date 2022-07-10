package main

import "fmt"

// https://leetcode.cn/problems/minimum-depth-of-binary-tree
func run111() {
	t := buildBinaryTreeByLevel([]int{2, null, 3, null, 4, null, 5, null, 6})
	fmt.Println(minDepth(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left != nil && root.Right != nil {
		l := minDepth(root.Left)
		r := minDepth(root.Right)

		if l > r {
			return r + 1
		} else {
			return l + 1
		}
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	return 1
}
