package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/subsets-ii/
func run90() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	fmt.Println(subsetsWithDup([]int{0}))
}

func subsetsWithDup(nums []int) [][]int {
	ans := make([][]int, 0)
	cur := make([]int, 0)
	visited := make(map[int]bool, len(nums)) //key: index of nums, value: visted or not
	sort.Ints(nums)

	var dfs func(int)
	dfs = func(k int) {
		if k == len(nums) {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		// case1: not add current number
		dfs(k + 1)

		// skip duplicated, skip current when it is same with prior and prior is not added
		if k > 0 && !visited[k-1] && nums[k-1] == nums[k] {
			return
		}

		// case1: add current number
		visited[k] = true
		cur = append(cur, nums[k])
		dfs(k + 1)
		visited[k] = false
		cur = cur[:len(cur)-1]

	}
	dfs(0)
	return ans
}
