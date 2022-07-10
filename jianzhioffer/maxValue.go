package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/li-wu-de-zui-da-jie-zhi-lcof/
func Run47() {
	fmt.Println(maxValue([][]int{{1, 3, 1}}))
}

func maxValue(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	ans := make([][]int, m)
	for i := 0; i < m; i++ {
		ans[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				ans[i][j] = grid[i][j]
			} else if i == 0 {
				ans[i][j] = grid[i][j] + ans[i][j-1]
			} else if j == 0 {
				ans[i][j] = grid[i][j] + ans[i-1][j]
			} else {
				max := ans[i][j-1]
				if ans[i-1][j] > max {
					max = ans[i-1][j]
				}
				ans[i][j] = grid[i][j] + max
			}
		}
	}
	return ans[m-1][n-1]
}
