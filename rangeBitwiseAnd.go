package main

import "fmt"

// https://leetcode.cn/problems/bitwise-and-of-numbers-range/
func run201() {
	fmt.Println(rangeBitwiseAnd(5, 7))
}

func rangeBitwiseAnd(left int, right int) int {
	moved := 0
	for left != right {
		left >>= 1
		right >>= 1
		moved++
	}
	return left << moved
}
