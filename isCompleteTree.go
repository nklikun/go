package main

import "fmt"

// https://leetcode.cn/problems/check-completeness-of-a-binary-tree
func run958() {
	// t := buildBinaryTreeByLevel([]int{1, 2, 3, 4, 5, null, 7, 8, 9, 10, 11, 12, 13})
	// t := buildBinaryTreeByLevel([]int{1, 2, 3, 5, null, 7, 8})
	t := buildBinaryTreeByLevel([]int{1, 2, 3, 5, null, 7, 8})
	fmt.Println(isCompleteTree(t))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNodei
 *     Right *TreeNode
 * }
 */
// func isCompleteTree(root *TreeNode) bool {
// 	//n是节点数，p是最大节点编号值
// 	n, p := 0, 0
// 	var helper func(*TreeNode, int) bool
// 	helper = func(node *TreeNode, k int) bool {
// 		//k是当前节点编号
// 		if node == nil {
// 			return true
// 		}
// 		if k > 100 {、
// 			//节点总数在100以内
// 			return false
// 		}
// 		n++
// 		if p < k {
// 			p = k
// 		}
// 		return helper(node.Left, 2*k) && helper(node.Right, 2*k+1)
// 	}
// 	if !helper(root, 1) {
// 		return false
// 	}
// 	return n == p
// }

// 层序遍历，nil后如果下一层或者右侧出现非nil，则true
func isCompleteTree(root *TreeNode) bool {
	// write code here
	cur := make([]*TreeNode, 0)
	cur = append(cur, root)
	var next []*TreeNode
	for len(cur) > 0 {
		next = make([]*TreeNode, 0)
		for i := 0; i < len(cur); i++ {
			if cur[i] == nil {
				for j := i + 1; j < len(cur); j++ {
					if cur[j] != nil {
						return false
					}
				}
				if len(next) > 0 {
					for j := 0; j < len(next); j++ {
						if next[j] != nil {
							return false
						}
					}
				}
			} else {
				next = append(next, cur[i].Left, cur[i].Right)
			}
		}
		cur = next
	}
	return true
}
