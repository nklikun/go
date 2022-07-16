package main

import "fmt"

/*
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
求在该柱状图中，能够勾勒出来的矩形的最大面积。

示例 1:

输入：heights = [2,1,5,6,2,3]
输出：10
解释：最大的矩形为图中红色区域，面积为 10
示例 2：

输入： heights = [2,4]
输出： 4

提示：
1 <= heights.length <=105
0 <= heights[i] <= 104


来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/largest-rectangle-in-histogram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run84() {
	fmt.Println(largestRectangleArea([]int{4, 2, 0, 3, 2, 4, 3, 4}))
	fmt.Println(largestRectangleArea([]int{2, 1, 5, 6, 2, 3}))
	fmt.Println(largestRectangleArea([]int{6, 7, 5, 2, 4, 5, 9, 3}))
}

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

/*
// 优化，使用单调递增栈代替左右查看的过程，并存储到left，right数组
func largestRectangleArea(heights []int) int {
	l, r := make([]int, len(heights)), make([]int, len(heights))
	stack := make([]int, 0)
	for i := 0; i < len(heights); i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			l[i] = -1
		} else {
			l[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	stack = make([]int, 0)
	for i := len(heights) - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			r[i] = len(heights)
		} else {
			r[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		area := (r[i] - l[i] - 1) * heights[i]
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
*/

/*// 暴力解（优化：减少重复计算）
func largestRectangleArea(heights []int) int {
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		l, r, h := i, i, heights[i]
		if h == 0 {
			continue
		}
		// 左边看一下，看最多能向左延伸多长，找到大于等于当前柱形高度的最左边元素的下标；
		// 右边看一下，看最多能向右延伸多长；找到大于等于当前柱形高度的最右边元素的下标。
		//
		// 为什么只找到大于等于当前高度的柱为止呢，因为小于当前柱体高度（假设为j）的面积
		// 自然在遍历到j时会计算得上，现在计算则会重复计算。
		for l >= 0 && r < len(heights) && heights[l] >= h && heights[r] >= h {
			maxArea = max(maxArea, (r-l+1)*h)
			if l == 0 {
				r++
			} else if r == len(heights)-1 {
				l--
			} else {
				// 往高的方向发展，否则可能低的部分可能拉低高的部分的面积值，遗漏高的部分
				if heights[l-1] < heights[r+1] {
					r++
				} else if heights[l-1] > heights[r+1] {
					l--
				} else {
					l--
					r++
				}
			}
		}
	}
	return maxArea
}
*/

/* // 暴力解(包含全部重复的计算)
func largestRectangleArea(heights []int) int {
	min := func(a int, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a int, b int) int {
		if a > b {
			return a
		}
		return b
	}
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		l, r, h := i, i, heights[i]
		for l >= 0 && r < len(heights) {
			h = min(h, min(heights[l], heights[r]))
			if h == 0 {
				break
			}
			maxArea = max(maxArea, (r-l+1)*h)

			if l == 0 {
				r++
			} else if r == len(heights)-1 {
				l--
			} else {
				// 往高的方向发展，否则可能低的部分可能拉低高的部分的面积值，遗漏高的部分
				if heights[l-1] < heights[r+1] {
					r++
				} else if heights[l-1] > heights[r+1] {
					l--
				} else {
					l--
					r++
				}
			}
		}
	}
	return maxArea
}
*/
