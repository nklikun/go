package main

// https://leetcode.cn/problems/search-insert-position/
func run35() {

}

func searchInsert(nums []int, target int) int {
	l, r, res := 0, len(nums)-1, len(nums)
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] >= target {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
