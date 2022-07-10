package jianzhioffer

// https://leetcode.cn/problems/zhong-jian-er-cha-shu-lcof/
func Run07() {
	buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
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
	if len(inorder) == 0 {
		return nil
	}

	preIndex, inIndex := 0, 0
	for i, pre := range preorder {
		for j, in := range inorder {
			if pre == in {
				preIndex = i
				inIndex = j
				break
			}
		}
	}
	node := &TreeNode{Val: inorder[inIndex]}
	node.Left = buildTree(preorder[preIndex+1:], inorder[0:inIndex])
	node.Right = buildTree(preorder[preIndex+1:], inorder[inIndex+1:])
	return node
}
