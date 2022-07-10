package main

// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/
func run80() {

}

func removeDuplicates2(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}

	l := 1
	for i := 2; i < len(nums); i++ {
		if nums[i] == nums[l] && nums[l] == nums[l-1] {
			continue
		}
		l++
		nums[l] = nums[i]
	}
	return l + 1
}
