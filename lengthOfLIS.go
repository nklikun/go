package main

import "fmt"

// https://leetcode.cn/problems/longest-increasing-subsequence/
func run300() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	max := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = lengthOfLIS_max(dp[j]+1, dp[i])
			}
		}
		max = lengthOfLIS_max(dp[i], max)
	}
	return max
}

func lengthOfLIS_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
