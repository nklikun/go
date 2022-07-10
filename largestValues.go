package main

// https://leetcode.cn/problems/find-largest-value-in-each-tree-row
func run515() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var next, cur []*TreeNode
	nums := make([]int, 0)
	cur = append(cur, root)
	for len(cur) > 0 {
		max := ^int(^uint32(0) >> 1)
		for _, node := range cur {
			if max < node.Val {
				max = node.Val
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		nums = append(nums, max)
		cur = next
		next = nil
	}
	return nums
}
