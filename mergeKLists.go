package main

import "fmt"

// https://leetcode.cn/problems/merge-k-sorted-lists
func run23() {
	l1 := buildNodeList([]int{1, 1, 4, 7})
	l2 := buildNodeList([]int{2, 3, 5, 7, 8})
	l3 := buildNodeList([]int{3, 3, 6, 9})

	l := mergeKLists([]*ListNode{l1, l2, l3})
	for node := l; node != nil; node = node.Next {
		fmt.Println(node.Val)
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	} else if len(lists) == 1 {
		return lists[0]
	} else if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	} else {
		half := int(len(lists) / 2)
		return mergeTwoLists(mergeKLists(lists[0:half]), mergeKLists(lists[half:]))
	}
}
