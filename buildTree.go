package main

import "fmt"

// https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal
func run105() {
	t := buildBinaryTreeByLevel([]int{3, 9, 20, NULL, NULL, 15, 7})
	preorder := preorderTraversal(t)
	inorder := inorderTraversal(t)
	fmt.Println(preorder)
	fmt.Println(inorder)
	t1 := buildTree(preorder, inorder)
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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	//找到第一个在inorder中存在的preorder元素，它即是当前root
	var node *TreeNode
	for i := 0; i < len(preorder); i++ {
		for j := 0; j < len(inorder); j++ {
			if preorder[i] == inorder[j] {
				node = &TreeNode{Val: preorder[i]}
				if j != 0 {
					node.Left = buildTree(preorder[i+1:], inorder[0:j])
				}
				if j != len(preorder)-1 {
					node.Right = buildTree(preorder[i+1:], inorder[j+1:])
				}
				return node
			}
		}
	}

	return node
}
