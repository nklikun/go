package main

// https://leetcode.cn/problems/remove-element
func run27() {

}

func removeElement(nums []int, val int) int {
	removed := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			removed++
		} else {
			nums[i-removed] = nums[i]
		}
	}
	return len(nums) - removed
}
