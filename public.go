package main

import "fmt"

const NULL int = -9999
const null int = -9999

//Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildNodeList(list []int) *ListNode {
	var head, cur *ListNode
	for _, i := range list {
		node := &ListNode{Val: i}
		if cur == nil {
			cur = node
			head = cur
		} else {
			cur.Next = node
			cur = cur.Next
		}
	}
	return head
}

func printNodeList(list *ListNode) {
	for node := list; node != nil; node = node.Next {
		fmt.Printf("%d  ", node.Val)
	}
	fmt.Println()
}

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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

func printBinaryTreeByLevel(root *TreeNode) {
	var cur, next []*TreeNode
	cur = append(cur, root)
	var nums []int
	for len(cur) != 0 {
		for i := 0; i < len(cur); i++ {
			node := cur[i]
			if node != nil {
				nums = append(nums, node.Val)
				next = append(next, node.Left, node.Right)
			} else {
				nums = append(nums, NULL)
			}
		}

		cur = next
		next = []*TreeNode{}
	}

	s := 0
	e := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != NULL {
			break
		}
		s++
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] != NULL {
			break
		}
		e++
	}
	for i := s; i < len(nums)-e; i++ {
		if nums[i] == NULL {
			fmt.Print("null")
		} else {
			fmt.Print(nums[i])
		}
		if i != len(nums)-e-1 {
			fmt.Print(",")
		}
	}
}

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
