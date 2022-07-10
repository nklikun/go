package main

import "fmt"

// https://leetcode.cn/problems/kth-smallest-element-in-a-bst
func run230() {
	t := buildBinaryTreeByLevel([]int{3, 1, 4, null, 2})
	fmt.Println(kthSmallest(t, 1))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var inorder func(*TreeNode)
	num := 0
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		k--
		if k == 0 {
			num = node.Val
		}
		inorder(node.Right)
	}
	inorder(root)
	return num
}
