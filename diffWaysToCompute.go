package main

import (
	"fmt"
	"strconv"
)

// https://leetcode.cn/problems/different-ways-to-add-parentheses/
func run241() {
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}

func diffWaysToCompute(expression string) []int {
	if num, err := strconv.Atoi(expression); err == nil {
		return []int{num}
	}

	ans := make([]int, 0)
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' || expression[i] == '*' {
			left := diffWaysToCompute(expression[0:i])
			right := diffWaysToCompute(expression[i+1:])

			for _, l := range left {
				for _, r := range right {
					if expression[i] == '+' {
						ans = append(ans, l+r)
					} else if expression[i] == '-' {
						ans = append(ans, l-r)
					} else {
						ans = append(ans, l*r)
					}
				}
			}
		}
	}
	return ans
}
