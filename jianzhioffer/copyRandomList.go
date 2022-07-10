package jianzhioffer

// https://leetcode.cn/problems/fu-za-lian-biao-de-fu-zhi-lcof/
func Run35() {}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := make(map[*Node]*Node, 0)
	copiedHead := &Node{Val: head.Val}
	nodeMap[head] = copiedHead

	node := head
	copiedNode := copiedHead
	for node.Next != nil {
		copiedNext := &Node{Val: node.Next.Val}
		copiedNode.Next = copiedNext
		node = node.Next
		copiedNode = copiedNode.Next
		nodeMap[node] = copiedNode
	}

	copiedNode = copiedHead
	for node = head; node != nil; node = node.Next {
		random, _ := nodeMap[node.Random]
		copiedNode.Random = random
		copiedNode = copiedNode.Next
	}
	return copiedHead
}
