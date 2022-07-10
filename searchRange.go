package main

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func run34() {}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target && nums[r] == target {
			return []int{l, r}
		}
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			if nums[l] != target {
				l = mid
				for ; nums[l] == target; l-- {
				}
				l++
			} else {
				r = mid
				for ; nums[r] == target; r++ {
				}
				r--
			}
		}
	}
	return []int{-1, -1}
}
