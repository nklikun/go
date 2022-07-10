package main

func run99() {
	// t := buildBinaryTreeByLevel([]int{3, 1, 4, null, null, 2})
	t := buildBinaryTreeByLevel([]int{1, 3, null, null, 2})
	// recoverTree(t)
	printBinaryTreeByLevel(t)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*func recoverTree(root *TreeNode) {
	recoverSearchTree(root, nil)
}

func recoverSearchTree(root *TreeNode, maxNode *TreeNode) {
	if root == nil {
		return
	}

	recoverSearchTree(root.Left, maxNode)
	if maxNode != nil && root.Val < maxNode.Val {
		tmp := maxNode.Val
		maxNode.Val = root.Val
		root.Val = tmp
		return
	}
	maxNode = root
	recoverSearchTree(root.Right, maxNode)
	return
}*/
