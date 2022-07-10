package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/que-shi-de-shu-zi-lcof/
func Run53_2() {
	fmt.Println(missingNumber([]int{0, 1, 3, 4, 5, 6, 7, 8, 9}))
}

func missingNumber(nums []int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == mid {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}
