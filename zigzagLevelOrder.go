package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-zigzag-level-order-traversal
func run103() {
	t := buildBinaryTreeByLevel([]int{1, 2, 3, 4, null, null, 5})
	fmt.Println(zigzagLevelOrder(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	numsList := make([][]int, 0)
	cur := make([]*TreeNode, 0)
	next := make([]*TreeNode, 0)
	cur = append(cur, root)
	reverse := false
	for len(cur) > 0 {
		nums := make([]int, 0)
		for i := 0; i < len(cur); i++ {
			var node *TreeNode
			if reverse {
				node = cur[len(cur)-i-1]
			} else {
				node = cur[i]
			}
			nums = append(nums, node.Val)

			node = cur[i]
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
		reverse = !reverse
	}

	return numsList
}
