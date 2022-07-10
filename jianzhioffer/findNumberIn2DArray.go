package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/er-wei-shu-zu-zhong-de-cha-zhao-lcof/
func Run04() {
	fmt.Println(findNumberIn2DArray([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 15))
}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	m := len(matrix)
	n := len(matrix[0])
	if matrix[m-1][n-1] < target || matrix[0][0] > target {
		return false
	}

	for i, j := 0, n-1; i < m && j >= 0; {
		if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		} else {
			return true
		}
	}

	return false
}
