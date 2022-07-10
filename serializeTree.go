package main

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/serialize-and-deserialize-binary-tree
func run297() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
}

// func Constructor() Codec {
// 	var codec Codec
// 	return Codec
// }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var cur, next []*TreeNode
	cur = append(cur, root)
	var nums []string
	for len(cur) != 0 {
		for i := 0; i < len(cur); i++ {
			node := cur[i]
			if node != nil {
				nums = append(nums, strconv.Itoa(node.Val))
				next = append(next, node.Left, node.Right)
			} else {
				nums = append(nums, "nil")
			}
		}

		cur = next
		next = []*TreeNode{}
	}
	return strings.Join(nums, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	nums := strings.Split(data, ",")
	if len(nums) < 0 {
		return nil
	}

	if nums[0] == "nil" {
		return nil
	}
	val, _ := strconv.Atoi(nums[0])
	root := &TreeNode{Val: val}
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
			} else if nums[i] == "nil" {
				node.Left = nil
			} else {
				val, _ := strconv.Atoi(nums[i])
				left := &TreeNode{Val: val}
				node.Left = left
				nextLevel = append(nextLevel, left)
			}
			i++

			if i > len(nums)-1 {
				break
			} else if nums[i] == "nil" {
				node.Right = nil
				i++
				continue
			} else {
				val, _ := strconv.Atoi(nums[i])
				right := &TreeNode{Val: val}
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

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
