package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/zuo-xuan-zhuan-zi-fu-chuan-lcof/
func Run58() {}

func reverseLeftWords(s string, n int) string {
	if n >= len(s) || n <= 0 {
		return s
	}

	return fmt.Sprintf("%s%s", s[n:], s[:n])
}
