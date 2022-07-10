package main

import "fmt"

// https://leetcode.cn/problems/path-sum-ii
func run113() {
	t := buildBinaryTreeByLevel([]int{-6, null, -3, -6, 0, -6, -5, 4, null, null, null, 1, 7})
	fmt.Println(pathSum(t, -21))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, targetSum int) [][]int {
	var numsList [][]int

	var pathSumHelper func(*TreeNode, int, []int)
	pathSumHelper = func(root *TreeNode, targetSum int, nums []int) {
		if root == nil {
			return
		}

		if root.Left == nil && root.Right == nil {
			if targetSum == root.Val {
				nums = append(nums, root.Val)
				// need a deep copy of nums
				numsList = append(numsList, append([]int{}, nums...))
			}
			return
		}

		nums = append(nums, root.Val)
		pathSumHelper(root.Left, targetSum-root.Val, nums)
		pathSumHelper(root.Right, targetSum-root.Val, nums)
	}

	pathSumHelper(root, targetSum, []int{})

	return numsList
}
