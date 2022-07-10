package main

// https://leetcode.cn/problems/house-robber-iii
func run337() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rob(root *TreeNode) int {
	var rob_helper func(*TreeNode) (int, int)
	// return the values of node added or not added
	// return param 1:added, return param 2:not added
	rob_helper = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		l1, l2 := rob_helper(node.Left)
		r1, r2 := rob_helper(node.Right)
		n1 := node.Val + l2 + r2
		n2 := max(l1, l2) + max(r1, r2)
		return n1, n2
	}

	return max(rob_helper(root))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
