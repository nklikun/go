package jianzhioffer

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/
func Run32_1() {}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-ii-lcof/
func Run32_2() {}

// https://leetcode.cn/problems/cong-shang-dao-xia-da-yin-er-cha-shu-iii-lcof/
func Run32_3() {}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	res := make([]int, 0)
	cur, next := []*TreeNode{root}, make([]*TreeNode, 0)
	for len(cur) > 0 {
		for _, node := range cur {
			res = append(res, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		cur = next
		next = make([]*TreeNode, 0)
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	cur, next := []*TreeNode{root}, make([]*TreeNode, 0)
	for len(cur) > 0 {
		level := make([]int, 0)
		for _, node := range cur {
			level = append(level, node.Val)
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, level)
		cur = next
		next = make([]*TreeNode, 0)
	}
	return res
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder3(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	cur, next := []*TreeNode{root}, make([]*TreeNode, 0)
	reverse := false
	for len(cur) > 0 {
		level := make([]int, 0)
		for _, node := range cur {
			if reverse {
				level = append([]int{node.Val}, level...)
			} else {
				level = append(level, node.Val)
			}
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		res = append(res, level)
		cur = next
		next = make([]*TreeNode, 0)
		reverse = !reverse
	}
	return res
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
