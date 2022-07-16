package main

import (
	"fmt"
)

/*
给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
1 0 1 0 0
1 0 1 1 1
1 1 1 1 1
1 0 0 1 0

示例 1：

输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：6
解释：最大矩形如上图所示。
示例 2：

输入：matrix = []
输出：0
示例 3：

输入：matrix = [["0"]]
输出：0
示例 4：

输入：matrix = [["1"]]
输出：1
示例 5：

输入：matrix = [["0","0"]]
输出：0


提示：

rows == matrix.length
cols == matrix[0].length
1 <= row, cols <= 200
matrix[i][j] 为 '0' 或 '1'

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run85() {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'}}
	fmt.Println(maximalRectangle(matrix))
}

func maximalRectangle(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	maxRec := 0
	//逐行放入到heights
	heights := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > maxRec {
			maxRec = area
		}
	}
	return maxRec
}

/*
// 优化，使用单调递增栈代替左右查看的过程，并存储到left，right数组
func largestRectangleArea(heights []int) int {
	n := len(heights)
	l, r := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = n
	}
	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		sl := len(stack)
		for sl > 0 && heights[stack[sl-1]] >= heights[i] {
			r[stack[sl-1]] = i
			stack = stack[:sl-1]
			sl = len(stack)
		}
		if sl == 0 {
			l[i] = -1
		} else {
			l[i] = stack[sl-1]
		}
		stack = append(stack, i)
	}

	maxArea := 0
	for i := 0; i < n; i++ {
		area := (r[i] - l[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
*/
