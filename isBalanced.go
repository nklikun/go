package main

import "fmt"

// https://leetcode.cn/problems/balanced-binary-tree
func run110() {
	// t := buildBinaryTreeByLevel([]int{3, 9, 20, null, null, 15, 7})
	t := buildBinaryTreeByLevel([]int{1, 2, 3, 4, 5, 7, null, null, null, 6})
	fmt.Println(isBalanced(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {

	if root == nil {
		return true
	}

	l := highestOfTree(root.Left)
	r := highestOfTree(root.Right)

	return l-r < 2 && l-r > -2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func highestOfTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	l := highestOfTree(root.Left)
	r := highestOfTree(root.Right)
	if l > r {
		return l + 1
	}
	return r + 1
}
