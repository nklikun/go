package main

import (
	"fmt"
)

/*
209. 长度最小的子数组
给定一个含有 n 个正整数的数组和一个正整数 target 。
找出该数组中满足其和 ≥ target 的长度最小的 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，并返回其长度。如果不存在符合条件的子数组，返回 0 。

示例 1：
输入：target = 7, nums = [2,3,1,2,4,3]
输出：2
解释：子数组 [4,3] 是该条件下的长度最小的子数组。

示例 2：
输入：target = 4, nums = [1,4,4]
输出：1

示例 3：
输入：target = 11, nums = [1,1,1,1,1,1,1,1]
输出：0

提示：
1 <= target <= 109
1 <= nums.length <= 105
1 <= nums[i] <= 105

进阶：
如果你已经实现 O(n) 时间复杂度的解法, 请尝试设计一个 O(n log(n)) 时间复杂度的解法。
https://leetcode.cn/problems/minimum-size-subarray-sum/
*/

func run209() {
	fmt.Println(minSubArrayLen(5, []int{1, 4, 4}))                 // 2
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))        // 2
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1})) // 0

	fmt.Println(minSubArrayLen(11, []int{1, 2, 3, 4, 5}))                 // 3
	fmt.Println(minSubArrayLen(15, []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8})) // 2
}

// O(n)
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r, sum, max := 0, 0, 0, len(nums)+1
	for r < len(nums) {
		sum += nums[r]
		for sum >= target {
			if r-l+1 < max {
				max = r - l + 1
			}
			sum -= nums[l]
			l++
		}
		r++
	}
	if max == len(nums)+1 {
		return 0
	}
	return max
}

/*
// O(nlog(n))
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	sum := make([]int, n+1) // sum[i] = nums[0] + ... + nums[i-1]
	sum[0] = nums[0]        // 首位存0, 方便后续查找值为target - sum[i]的连续序列
	for i := 1; i <= n; i++ {
		sum[i] = sum[i-1] + nums[i-1]
	}

	if sum[n] < target {
		return 0
	}

	max := n + 1
	for i := 0; i < n; i++ {
		l, r, mid := i, n, 0
		for l < r {
			mid = (l + r) / 2
			if sum[mid] < target+sum[i] {
				l = mid + 1
			} else {
				r = mid
			}
		}
		if sum[l]-sum[i] >= target && max > (l-i) {
			max = l - i
		}
	}

	if max > n {
		return 0
	}
	return max
}
*/
