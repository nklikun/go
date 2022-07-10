package main

import "fmt"

// https://leetcode.cn/problems/validate-binary-search-tree
func run98() {
	// t := buildBinaryTreeByLevel([]int{5, 3, 7, 1, 4, 6, 8, 0, 2, NULL, NULL, NULL, NULL, NULL, 9})
	t := buildBinaryTreeByLevel([]int{1, 1})
	fmt.Println(isValidBST(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	var MAX int = ^(int(^uint32(0) >> 1))
	return isValidBST1(root, &MAX)
}

func isValidBST1(root *TreeNode, max *int) bool {
	if root == nil {
		return true
	}

	if isValidBST1(root.Left, max) {
		if *max < root.Val {
			*max = root.Val
			return isValidBST1(root.Right, max)
		}
	}

	return false
}
