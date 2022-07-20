package main

import (
	"fmt"
	"sort"
)

/*
611. 有效三角形的个数
给定一个包含非负整数的数组 nums ，返回其中可以组成三角形三条边的三元组个数。

示例 1:
输入: nums = [2,2,3,4]
输出: 3
解释:有效的组合是:
2,3,4 (使用第一个 2)
2,3,4 (使用第二个 2)
2,2,3

示例 2:
输入: nums = [4,2,3,4]
输出: 4

提示:
1 <= nums.length <= 1000
0 <= nums[i] <= 1000

https://leetcode.cn/problems/valid-triangle-number/
*/

func run611() {
	fmt.Println(triangleNumber([]int{2, 2, 3, 4})) // 3
	fmt.Println(triangleNumber([]int{4, 2, 3, 4})) // 4
}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	total := 0
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			sumTwo := nums[i] + nums[j]
			l, r, k := j+1, len(nums)-1, j
			for l <= r {
				mid := (l + r) / 2
				if nums[mid] >= sumTwo {
					r = mid - 1
				} else {
					k = mid
					l = mid + 1
				}
			}
			total += k - (j + 1) + 1
		}
	}
	return total
}
