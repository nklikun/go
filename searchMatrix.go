package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/search-a-2d-matrix/
func run74() {
	b := searchMatrix([][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
		{10, 11, 12, 13, 15, 16, 17, 18, 19},
		{20, 21, 22, 23, 25, 26, 27, 28, 29},
		{30, 31, 32, 33, 34, 35, 37, 38, 39},
		{40, 42, 43, 44, 45, 46, 47, 48, 49},
		{50, 51, 52, 53, 54, 55, 56, 57, 58}}, 42)
	// b := searchMatrix([][]int{{1, 3}}, 3)
	fmt.Println(b)
}

func searchMatrix(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool { return matrix[i][0] > target }) - 1
	if row < 0 {
		return false
	}
	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}
