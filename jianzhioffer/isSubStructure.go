package jianzhioffer

// https://leetcode.cn/problems/shu-de-zi-jie-gou-lcof/
func Run26() {}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	if A.Val == B.Val {
		if isSubStructure_helper(A, B) {
			return true
		}
	}
	return isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func isSubStructure_helper(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return isSubStructure_helper(a.Left, b.Left) && isSubStructure_helper(a.Right, b.Right)
}
