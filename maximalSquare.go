package main

import "fmt"

// https://leetcode.cn/problems/maximal-square/
func run221() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalSquare(matrix))
}

func maximalSquare(matrix [][]byte) int {
	h := len(matrix)
	l := len(matrix[0])
	dps := make([][]int, h)
	max := 0
	for i := 0; i < h; i++ {
		dps[i] = make([]int, l)
		for j := 0; j < l; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dps[i][j] = 1
					max = maximalSquare_max(max, dps[i][j])
				} else {
					dps[i][j] = 1 + maximalSquare_min(dps[i-1][j], dps[i][j-1], dps[i-1][j-1])
					max = maximalSquare_max(max, dps[i][j])
				}
			} else {
				dps[i][j] = 0
			}
		}
	}
	return max * max
}

func maximalSquare_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalSquare_min(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	min := nums[0]
	for _, n := range nums {
		if min > n {
			min = n
		}
	}
	return min
}
