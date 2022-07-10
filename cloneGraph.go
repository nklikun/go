package main

// https://leetcode.cn/problems/clone-graph
func run133() {
	n1 := &GraphNode{Val: 1}
	n2 := &GraphNode{Val: 2}
	n3 := &GraphNode{Val: 3}
	n1.Neighbors = []*GraphNode{n2, n3}
	n2.Neighbors = []*GraphNode{n1, n3}
	n3.Neighbors = []*GraphNode{n1, n2}
	cloneGraph(n1)
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *GraphNode) *GraphNode {
	var cloneGraph_helper func(*GraphNode) *GraphNode
	clonedNodes := make([]*GraphNode, 0)
	cloneGraph_helper = func(node *GraphNode) *GraphNode {
		if node == nil {
			return nil
		}

		for _, ele := range clonedNodes {
			if ele.Val == node.Val {
				return ele
			}
		}

		newNode := &GraphNode{Val: node.Val}
		clonedNodes = append(clonedNodes, newNode)
		if node.Neighbors == nil {
			newNode.Neighbors = nil
			return newNode
		}
		neighbors := make([]*GraphNode, 0)
		for _, n := range node.Neighbors {
			clonedN := cloneGraph_helper(n)
			neighbors = append(neighbors, clonedN)
		}
		newNode.Neighbors = neighbors
		return newNode
	}
	return cloneGraph_helper(node)
}

// Definition for a Node.
type GraphNode struct {
	Val       int
	Neighbors []*GraphNode
}

/*

func cloneGraph(node *Node) *Node {
	var cloneGraph_helper func(*Node) *Node
	clonedNodes := make([]*Node, 0)
	cloneGraph_helper = func(node *Node) *Node {
		if node == nil {
			return nil
		}

		for _, ele := range clonedNodes {
			if ele.Val == node.Val {
				return ele
			}
		}

		newNode := &Node{Val: node.Val}
		clonedNodes = append(clonedNodes, newNode)
		if node.Neighbors == nil {
			newNode.Neighbors = nil
			return newNode
		}
		neighbors := make([]*Node, 0)
		for _, n := range node.Neighbors {
			clonedN := cloneGraph_helper(n)
			neighbors = append(neighbors, clonedN)
		}
		newNode.Neighbors = neighbors
		return newNode
	}
	return cloneGraph_helper(node)
}
*/
