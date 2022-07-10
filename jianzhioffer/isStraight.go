package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/bu-ke-pai-zhong-de-shun-zi-lcof/
func Run61() {
	fmt.Println(isStraight([]int{11, 8, 12, 8, 10}))
}

func isStraight(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	k := 0
	for ; k < len(nums) && nums[k] == 0; k++ {
	}
	for i := k; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return false
		}
	}
	if k == 5 {
		return true
	}
	return nums[4]-nums[k] < 5
}
