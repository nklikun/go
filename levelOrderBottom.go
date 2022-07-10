package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-level-order-traversal-ii
func run107() {
	t := buildBinaryTreeByLevel([]int{3, 9, 20, null, null, 15, 7})
	fmt.Println(levelOrderBottom(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	var cur, next []*TreeNode
	cur = append(cur, root)
	var nums []int
	numslist := make([][]int, 0)
	for len(cur) != 0 {
		for i := 0; i < len(cur); i++ {
			node := cur[i]
			if node != nil {
				nums = append(nums, node.Val)
				next = append(next, node.Left, node.Right)
			}
		}

		cur = next
		if len(nums) > 0 {
			numslist = append([][]int{nums}, numslist...)
		}
		next = []*TreeNode{}
		nums = []int{}
	}
	return numslist
}
