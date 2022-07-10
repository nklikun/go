package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-level-order-traversal
func run102() {
	t := buildBinaryTreeByLevel([]int{3, 9, 20, NULL, NULL, 15, 7})
	fmt.Println(levelOrder(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	numsList := make([][]int, 0)
	cur := make([]*TreeNode, 0)
	next := make([]*TreeNode, 0)
	cur = append(cur, root)
	for len(cur) > 0 {
		nums := make([]int, 0)
		for _, node := range cur {
			nums = append(nums, node.Val)

			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		numsList = append(numsList, nums)
		cur = next
		next = []*TreeNode{}
	}

	return numsList
}
