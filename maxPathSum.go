package main

import "fmt"

// https://leetcode.cn/problems/binary-tree-maximum-path-sum
func run124() {
	t := buildBinaryTreeByLevel([]int{1, 2, 3})
	fmt.Println(maxPathSum(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxPathSum(root *TreeNode) int {
	max := ^int(^uint32(0) >> 1)

	var calNodeValue func(*TreeNode) int
	calNodeValue = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		l := calNodeValue(root.Left)
		r := calNodeValue(root.Right)

		// 返回更大的节点值作为父节点计算使用的子树值
		var ret int
		if l < 0 && r < 0 {
			ret = root.Val
		} else if l > r {
			ret = l + root.Val
		} else {
			ret = r + root.Val
		}

		// 如果root节点参与的计算能获取最大值，则记录最大值
		if ret > max {
			max = ret
		}
		//计算左右树和当前节点值能否得到最大值
		// 返回值不能把左右子树和根节点同时包含在内，所以此值单独计算不作为返回值的可能
		if l+r+root.Val > max {
			max = l + r + root.Val
		}

		return ret
	}
	calNodeValue(root)

	return max
}
