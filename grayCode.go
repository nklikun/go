package main

import "fmt"

// https://leetcode.cn/problems/gray-code/
func run89() {
	fmt.Println(grayCode(1))
}

func grayCode(n int) []int {
	ans := make([]int, 0)
	ans = append(ans, 0)
	for i := 1; i <= n; i++ {
		for j := len(ans) - 1; j >= 0; j-- {
			ans = append(ans, (1<<(i-1))+ans[j])
		}
	}
	return ans
}
