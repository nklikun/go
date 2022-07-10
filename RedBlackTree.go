package main

// https://leetcode.cn/circle/article/nsG69E/
func run_RBTree() {

}

type COLOR bool

const (
	Black COLOR = true
	Red   COLOR = false
)

type RBTreeNode struct {
	Val    int
	Color  COLOR
	Parent *RBTreeNode
	Left   *RBTreeNode
	Right  *RBTreeNode
}

func leftRotate(pNode *RBTreeNode) {
	if pNode == nil {
		return
	}
	vNode := pNode.Right
	rNode := pNode.Left
	pNode.Left = vNode
	if rNode != nil {
		rNode.Parent = pNode
	}

	vNode.Parent = pNode.Parent
	if pNode.Parent != nil {
		if pNode.Parent.Left == pNode {
			pNode.Parent.Left = vNode
		} else {
			pNode.Parent.Right = vNode
		}
	}

	vNode.Left = pNode
	pNode.Parent = vNode
}

func rightRotate(pNode *RBTreeNode) {
	if pNode == nil {
		return
	}
	fNode := pNode.Left
	kNode := pNode.Right
	pNode.Left = kNode

	if kNode != nil {
		kNode.Parent = pNode
	}

	fNode.Parent = pNode.Parent
	if pNode.Parent != nil {
		if pNode.Parent.Left == pNode {
			pNode.Parent.Left = fNode
		} else {
			pNode.Parent.Right = fNode
		}
	}

	fNode.Right = pNode
	pNode.Parent = fNode
}

func insert(root *RBTreeNode, nodeInserted *RBTreeNode) bool {
	if root == nil {
		root = nodeInserted
		return true
	}

	if root.Val > nodeInserted.Val {
		if root.Left == nil {
			root.Left = nodeInserted
			nodeInserted.Parent = root
			nodeInserted.Color = Red
			insertFixUp(nodeInserted)
			return true
		} else {
			return insert(root.Left, nodeInserted)
		}
	} else if root.Val < nodeInserted.Val {
		if root.Right == nil {
			root.Right = nodeInserted
			nodeInserted.Parent = root
			nodeInserted.Color = Red
			insertFixUp(nodeInserted)
			return true
		} else {
			return insert(root.Right, nodeInserted)
		}
	} else {
		return false
	}
}

func insertFixUp(node *RBTreeNode) {

}
