package main

// https://leetcode.cn/problems/binary-tree-tilt
func run563() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	ans := 0
	var sumOfTree func(*TreeNode) int
	sumOfTree = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l, r := sumOfTree(node.Left), sumOfTree(node.Right)
		if l > r {
			ans += l - r
		} else {
			ans += r - l
		}
		return node.Val + l + r
	}

	sumOfTree(root)
	return ans
}
