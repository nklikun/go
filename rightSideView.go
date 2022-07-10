package main

// https://leetcode.cn/problems/binary-tree-right-side-view
func run199() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	nums := make([]int, 0)
	var cur, next []*TreeNode
	cur = append(cur, root)
	for len(cur) > 0 {
		nums = append(nums, cur[len(cur)-1].Val)
		for i := 0; i < len(cur); i++ {
			if cur[i].Left != nil {
				next = append(next, cur[i].Left)
			}

			if cur[i].Right != nil {
				next = append(next, cur[i].Right)
			}
		}
		cur = next
		next = []*TreeNode{}
	}
	return nums
}
