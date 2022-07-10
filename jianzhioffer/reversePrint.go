package jianzhioffer

// https://leetcode.cn/problems/cong-wei-dao-tou-da-yin-lian-biao-lcof/
func Run06() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) []int {
	reverse := make([]int, 0)
	for ; head != nil; head = head.Next {
		reverse = append(reverse, head.Val)
	}
	nums := make([]int, len(reverse))
	for i, n := range reverse {
		nums[len(reverse)-1-i] = n
	}
	return nums
}
