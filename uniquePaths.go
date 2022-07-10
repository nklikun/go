package main

import "fmt"

// https://leetcode.cn/problems/unique-paths/
func run62() {
	fmt.Println(uniquePaths(3, 7))
}

func uniquePaths(m int, n int) int {
	paths := make([][]int, m)
	for i := 0; i < m; i++ {
		paths[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				paths[i][j] = 1
			} else {
				paths[i][j] = paths[i-1][j] + paths[i][j-1]
			}
		}
	}
	return paths[m-1][n-1]
}
