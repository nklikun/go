package main

// https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree
func run108() {
	t1 := sortedArrayToBST([]int{1, 2, 3, 4, 5, 6})
	printBinaryTreeByLevel(t1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := int(len(nums) / 2)

	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}
