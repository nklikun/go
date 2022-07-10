package main

import "fmt"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node
func run116() {
	t := run116_build([]int{1, 2, 3, 4, 5, 6, 7})
	n := connect(t)
	run116_print(n)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	// 各级别左子树最右节点指向右子树最左节点
	l := root.Left
	r := root.Right
	for l != nil {
		l.Next = r
		l = l.Right
		r = r.Left
	}

	connect(root.Left)
	connect(root.Right)

	return root
}

// NULL is the spcial number stand for NULL
func run116_build(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	root := &Node{Val: nums[0]}
	node := root
	level := make([]*Node, 0)
	nextLevel := make([]*Node, 0)
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

				left := &Node{Val: nums[i]}
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
				right := &Node{Val: nums[i]}
				node.Right = right
				nextLevel = append(nextLevel, right)
			}
			i++

		}

		level = nextLevel
		nextLevel = make([]*Node, 0)
	}
	return root
}

func run116_print(root *Node) {
	for root != nil {
		next := root.Left
		for root != nil {
			fmt.Print(root.Val)
			fmt.Print(",")
			root = root.Next
		}
		fmt.Print("#")
		root = next
	}
}
