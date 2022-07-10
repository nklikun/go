package main

import "fmt"

// https://leetcode.cn/problems/implement-strstr
func run28() {
	fmt.Println(strStr("aaa", "aaaaa"))
}

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		j := 0

		if len(haystack)-i < len(needle) {
			return -1
		}
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	return -1
}
