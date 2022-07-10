package main

import "fmt"

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array
func run26() {
	removeDuplicates([]int{})
}

func removeDuplicates(nums []int) int {
	removed := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-removed-1] {
			nums[i-removed] = nums[i]
		} else {
			removed++
		}
	}

	fmt.Println(nums)
	fmt.Println(len(nums) - removed)
	return len(nums) - removed
}
