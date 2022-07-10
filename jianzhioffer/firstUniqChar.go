package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/di-yi-ge-zhi-chu-xian-yi-ci-de-zi-fu-lcof/
func Run50() {
	fmt.Println(string(firstUniqChar("abaccdeff")))
}

func firstUniqChar(s string) byte {
	chars := make([]int, 26)
	for _, c := range s {
		chars[c-'a']++
	}

	for _, c := range s {
		if chars[c-'a'] == 1 {
			return byte(c)
		}
	}
	return ' '
}
