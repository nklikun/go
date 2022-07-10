package main

import "fmt"

// https://leetcode.cn/problems/validate-binary-tree-nodes/
func run1361() {
	// fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, -1, -1, -1}))
	// fmt.Println(validateBinaryTreeNodes(4, []int{1, -1, 3, -1}, []int{2, 3, -1, -1}))
	// fmt.Println(validateBinaryTreeNodes(2, []int{1, 0}, []int{-1, -1}))
	// fmt.Println(validateBinaryTreeNodes(6, []int{1, -1, -1, 4, -1, -1}, []int{2, -1, -1, 5, -1, -1}))
	// fmt.Println(validateBinaryTreeNodes(6, []int{1, 3, 5, -1, -1, -1}, []int{2, 4, -1, -1, -1, -1}))
	fmt.Println(validateBinaryTreeNodes(1, []int{-1}, []int{-1}))
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// 需要验证的内容: 1. 节点无环（不存在无前驱节点）
	// 2. 顶点只有一个（即初始入度不存在的只有节点0)，3. 任何节点父节点只有一个
	// 易判断：1. 节点入度为0的点只有一个且为root，2. 从root遍历，所有节点必须且只能被遍历一次

	// 入度，存储每个节点的入度,即父节点
	degree := make([]int, n)
	// key: node n, value: {0:left child, 1:right child}
	nodes := make(map[int][]int, n)
	for parent, child := range leftChild {
		nodes[parent] = []int{child, -1}
		if child != -1 {
			degree[child]++
		}
	}

	for parent, child := range rightChild {
		if nodes[parent] != nil {
			nodes[parent][1] = child
		} else {
			nodes[parent] = []int{-1, child}
		}
		if child != -1 {
			degree[child]++
		}
	}

	root := -1
	for i, p := range degree {
		if p == 0 {
			if root != -1 {
				return false
			}
			root = i
		}
	}
	if root == -1 {
		return false
	}

	cnt := 0
	var traverseTree func(int) bool
	traverseTree = func(node int) bool {
		// 借用degree来存储遍历过的节点
		if degree[node] == -2 {
			return false
		}
		degree[node] = -2
		cnt++
		left := nodes[node][0]
		right := nodes[node][1]
		l, r := true, true
		if left != -1 {
			l = traverseTree(left)
		}
		if right != -1 {
			r = traverseTree(right)
		}
		return l && r
	}

	return traverseTree(root) && cnt == n
}
