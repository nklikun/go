package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-paths
func run257() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	paths := make([]string, 0)
	var binaryTreePath func(*TreeNode, string)
	binaryTreePath = func(node *TreeNode, path string) {
		if node == nil {
			return
		}

		if node.Left == nil && node.Right == nil {
			path = path + fmt.Sprintf("%d", node.Val)
			paths = append(paths, path)
		}

		path = path + fmt.Sprintf("%d->", node.Val)
		if node.Left != nil {
			binaryTreePath(node.Left, path)
		}
		if node.Right != nil {
			binaryTreePath(node.Right, path)
		}
	}

	binaryTreePath(root, "")
	return paths
}
