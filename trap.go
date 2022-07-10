package main

import "fmt"

// https://leetcode.cn/problems/trapping-rain-water
func run42() {
	fmt.Print(trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}

func trap(height []int) int {
	l, r, lmax, rmax, ans := 0, len(height)-1, 0, 0, 0
	for l < r {
		lmax = trap_max(lmax, height[l])
		rmax = trap_max(rmax, height[r])
		if height[r] > height[l] {
			ans += lmax - height[l]
			l++
		} else {
			ans += rmax - height[r]
			r--
		}

	}
	return ans
}

func trap_max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
