package main

// https://leetcode.cn/problems/reverse-nodes-in-k-group
func run25() {
	list1 := buildNodeList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	printNodeList(list1)
	printNodeList(reverseKGroup(list1, 4))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	index := 0
	node, start := head, head
	var newList, newEnd *ListNode
	for {
		index++
		if index%k == 0 && start != nil {
			start = node.Next
			node.Next = nil

			tmpList := reverseList(head)

			if nil == newList {
				newList = tmpList
			} else {
				newEnd.Next = tmpList
			}
			// keep the end pointer of new list
			newEnd = head

			// reset head and node to next part
			head = start
			node = start
			continue
		}

		if node == nil || node.Next == nil {
			break
		}
		node = node.Next
	}

	if start != nil {
		newEnd.Next = start
	}

	return newList
}

func reverseList(head *ListNode) *ListNode {
	var first *ListNode
	second := head
	third := head.Next

	for {
		second.Next = first
		if third == nil {
			break
		}
		first = second
		second = third
		third = third.Next
	}
	return second
}
