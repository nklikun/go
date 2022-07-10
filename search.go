package main

import "fmt"

// https://leetcode.cn/problems/search-in-rotated-sorted-array
func run33() {
	nums := []int{5, 6, 7, 1, 2, 3, 4}
	b := search(nums, 3)
	fmt.Println(b)
}

func search(nums []int, target int) int {
	return searchIndexOfNums(0, len(nums)-1, nums, target)
}

func searchIndexOfNums(s, e int, nums []int, target int) int {
	if s == e {
		if nums[s] != target {
			return -1
		} else {
			return s
		}
	}

	n := int((e - s + 1) / 2)
	if (nums[s] <= nums[s+n-1] && target >= nums[s] && target <= nums[s+n-1]) ||
		(nums[s] > nums[s+n-1] && (target < nums[s+n] || target > nums[e])) {
		return searchIndexOfNums(s, s+n-1, nums, target)
	} else {
		return searchIndexOfNums(s+n, e, nums, target)
	}
}
