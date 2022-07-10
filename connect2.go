package main

import "fmt"

// https://leetcode.cn/problems/populating-next-right-pointers-in-each-node-ii
func run117() {
	/*
		/*
		                            1
								  /   \
								 2     3
								/\      \
							   /  \      \
							  4    5      6
							 /             \
							7               8
	*/
	t := run117_build([]int{1, 2, 3, 4, 5, null, 6, 7, null, null, null, null, 8})
	n := connect2(t)
	fmt.Println()
	run117_print(n)
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

func connect2(root *Node) *Node {
	if root == nil {
		return nil
	}

	var last *Node
	var linkToNext func(*Node)
	nextHead := root
	curHead := root

	linkToNext = func(node *Node) {
		if node == nil {
			return
		}
		if last == nil {
			last = node
			nextHead = node
			if node != nil {
				fmt.Print("\nnext: ")
				fmt.Print(node.Val)
				fmt.Print(",")
			}
		} else {
			last.Next = node
			last = last.Next
			if node != nil {
				fmt.Print(node.Val)
				fmt.Print(",")
			}
		}
	}

	for curHead != nil {
		for curHead != nil {
			linkToNext(curHead.Left)
			linkToNext(curHead.Right)
			curHead = curHead.Next
		}
		curHead = nextHead
		nextHead = nil
		last = nil
	}
	return root
}

// NULL is the spcial number stand for NULL
func run117_build(nums []int) *Node {
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

func run117_print(root *Node) {
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
