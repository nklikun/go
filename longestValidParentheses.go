package main

import "fmt"

// https://leetcode.cn/problems/longest-valid-parentheses
func run32() {
	s := "()()(())"
	fmt.Println(longestValidParentheses(s))
}

func longestValidParentheses(s string) int {
	length := len(s)
	if length == 0 {
		return 0
	}

	ss := s[:]
	left, max := 0, 0
	dp := make([]int, length)
	for i := 0; i < length; i++ {
		if ss[i] == '(' {
			// if i != 0 && left == 0 {
			// 	dp[i] = dp[i-1]
			// }
			left++
		} else {
			if left > 0 {
				left--
				// 完成一个阶段处理时，把上一阶段处理结果加上（可能是匹配完整的值，视作连续，也可能是没有匹配完整的0值）
				// 比如()(())处理到最后一个')'时，需要把第一个')'也加上
				prev := i - 1 - (dp[i-1] + 1)
				if prev < 0 {
					dp[i] = dp[i-1] + 2
				} else {
					dp[i] = dp[i-1] + 2 + dp[prev]
				}
				if max < dp[i] {
					max = dp[i]
				}
			} else {
				left = 0
				dp[i] = 0
			}
		}
	}

	return max
}
