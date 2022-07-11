package main

import "fmt"

// https://leetcode.cn/problems/subsets/
func run78() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	ans := make([][]int, 0)

	total := 1 << len(nums)
	for i := 0; i < total; i++ {
		cur := make([]int, 0)
		for j, num := range nums {
			if i&(1<<j) != 0 {
				cur = append(cur, num)
			}
		}
		ans = append(ans, cur)
	}
	return ans
}
