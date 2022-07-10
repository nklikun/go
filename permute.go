package main

import "fmt"

// https://leetcode.cn/problems/permutations/
func run46() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	data := make([]int, 0)
	var dfs func(int)
	dfs = func(k int) {
		if k == len(nums) {
			tmp := make([]int, len(data))
			copy(tmp, data)
			ans = append(ans, tmp)
			return
		}
		for i := 0; i <= len(data); i++ {
			// data = append(append(data[:i], nums[k]), data[i:]...)
			data = append(data[:i], append([]int{nums[k]}, data[i:]...)...)
			dfs(k + 1)
			data = append(data[:i], data[i+1:]...)
		}
	}
	data = append(data, nums[0])
	dfs(1)
	return ans
}
