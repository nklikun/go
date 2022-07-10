package main

import "fmt"

// https://leetcode.cn/problems/count-complete-tree-nodes
func run222() {
	t := buildBinaryTreeByLevel([]int{1, 2, 3, 4, 5, 6})
	fmt.Println(countNodes(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	h := heightOfCompleteTree(root)
	return getBottomCount(root) + (1<<(h-1) - 1)
}

func getBottomCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	bottomCount := 0
	hl := heightOfCompleteTree(root.Left)
	hr := heightOfCompleteTree(root.Right)

	if hl == hr && hl != 0 {
		bottomCount += 1 << (hl - 1)
		bottomCount += getBottomCount(root.Right)
	} else {
		bottomCount += getBottomCount(root.Left)
	}

	return bottomCount
}

func heightOfCompleteTree(root *TreeNode) int {
	height := 0
	for ; root != nil; root = root.Left {
		height++
	}
	return height
}
