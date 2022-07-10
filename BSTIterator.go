package main

// https://leetcode.cn/problems/binary-search-tree-iterator
func run173() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	nums []int
}

/*
func Constructor(root *TreeNode) BSTIterator {
	var it BSTIterator
	it.inorder(root)
	return it
}
*/

func (this *BSTIterator) inorder(root *TreeNode) {
	if root == nil {
		return
	}
	this.inorder(root.Left)
	this.nums = append(this.nums, root.Val)
	this.inorder(root.Right)
}

func (this *BSTIterator) Next() int {
	num := this.nums[0]
	this.nums = this.nums[1:]
	return num
}

func (this *BSTIterator) HasNext() bool {
	return len(this.nums) > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
