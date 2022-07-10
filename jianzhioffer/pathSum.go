package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/
func Run34() {
	t := buildBinaryTreeByLevel([]int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, 5, 1})
	fmt.Println(pathSum(t, 22))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, target int) [][]int {
	paths := make([][]int, 0)
	path := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(node *TreeNode, sum int) {
		if node == nil {
			return
		}

		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && sum+node.Val == target {
			tmp := make([]int, len(path))
			copy(tmp, path)
			paths = append(paths, tmp)
		} else {
			dfs(node.Left, sum+node.Val)
			dfs(node.Right, sum+node.Val)
		}
		// back trace
		path = path[:len(path)-1]
	}
	dfs(root, 0)
	return paths
}

const NULL int = -9999
const null int = -9999

// NULL is the spcial number stand for NULL
func buildBinaryTreeByLevel(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	node := root
	level := make([]*TreeNode, 0)
	nextLevel := make([]*TreeNode, 0)
	level = append(level, node)
	i := 1
	for {
		if i > len(nums)-1 {
			break
		}

		for _, node := range level {
			if i > len(nums)-1 {
				break
			} else if nums[i] == NULL {
				node.Left = nil
			} else {

				left := &TreeNode{Val: nums[i]}
				node.Left = left
				nextLevel = append(nextLevel, left)
			}
			i++

			if i > len(nums)-1 {
				break
			} else if nums[i] == NULL {
				node.Right = nil
				i++
				continue
			} else {
				right := &TreeNode{Val: nums[i]}
				node.Right = right
				nextLevel = append(nextLevel, right)
			}
			i++

		}

		level = nextLevel
		nextLevel = make([]*TreeNode, 0)
	}
	return root
}
