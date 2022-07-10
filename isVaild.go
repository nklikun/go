package main

import "strings"

// https://leetcode.cn/problems/valid-parentheses
func run20() {
	isValid("")
}

func isValid(s string) bool {
	len1 := len(s)
	len2 := 0
	for len1 != len2 {
		len1 = len(s)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
		len2 = len(s)
	}

	return len2 == 0
}
