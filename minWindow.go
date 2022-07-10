package main

import "fmt"

// https://leetcode.cn/problems/minimum-window-substring/
func run76() {
	fmt.Println(string(minWindow("a", "a")))
}

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	orignal := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		orignal[t[i]]++
	}

	windowMap := make(map[byte]int)
	minL, minR := -1, -1
	l, r := 0, 0

	var check func() bool
	check = func() bool {
		for k, v := range orignal {
			if windowMap[k] < v {
				return false
			}
		}
		return true
	}
	var updateLR func()
	updateLR = func() {
		if minL == -1 || (r-l) < (minR-minL) {
			minL = l
			minR = r
		}
	}

	for r < len(s) {
		windowMap[s[r]]++
		if !check() {
			r++
			continue
		}

		for l <= r {
			updateLR()
			windowMap[s[l]]--
			l++
			if !check() {
				r++
				break
			}
		}
	}
	if minL == -1 {
		return ""
	}
	return s[minL : minR+1]
}
