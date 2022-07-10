package main

import "fmt"

// https://leetcode.cn/problems/construct-binary-tree-from-inorder-and-postorder-traversal
func run106() {
	t := buildBinaryTreeByLevel([]int{3, 9, 20, NULL, NULL, 15, 7})
	postorder := postorderTraversal(t)
	inorder := inorderTraversal(t)
	fmt.Println(postorder)
	fmt.Println(inorder)
	t1 := buildTree2(inorder, postorder)
	fmt.Println(t1)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	lp := len(postorder)
	li := len(inorder)
	// 找到当前的根节点，后续遍历的最后一个节点
	for i := lp - 1; i >= 0; i-- {
		for j := 0; j < li; j++ {
			if postorder[i] == inorder[j] {
				node := &TreeNode{Val: postorder[i]}
				if j != 0 {
					node.Left = buildTree2(inorder[:j], postorder[:i])
				}
				if j != li-1 {
					node.Right = buildTree2(inorder[j+1:], postorder[0:i])
				}
				return node
			}
		}
	}
	return nil
}
