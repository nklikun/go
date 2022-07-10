package main

// https://leetcode.cn/problems/merge-two-sorted-lists
func run21() {
	l1 := buildNodeList([]int{1, 2, 4})
	l2 := buildNodeList([]int{1, 2, 5})
	mergeTwoLists(l1, l2)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var node, head *ListNode
	for list1 != nil || list2 != nil {
		add1 := true

		if list1 == nil {
			add1 = false
		} else if list2 == nil {
			add1 = true
		} else if list1.Val > list2.Val {
			add1 = false
		} else {
			add1 = true
		}

		tmp := list1
		if !add1 {
			tmp = list2
			list2 = list2.Next
		} else {
			list1 = list1.Next
		}

		if node == nil {
			node = tmp
			head = node
		} else {
			node.Next = tmp
			node = node.Next
		}
	}
	return head
}
