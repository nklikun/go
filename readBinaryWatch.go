package main

import "fmt"

// https://leetcode.cn/problems/binary-watch/
func run401() {}

func readBinaryWatch(turnedOn int) []string {
	ans := make([]string, 0)
	for i := 0; i < 12; i++ {
		for j := 0; j < 60; j++ {
			if getCountOfOne(i)+getCountOfOne(j) == turnedOn {
				var min string
				if j < 10 {
					min = fmt.Sprintf("0%d", j)
				} else {
					min = fmt.Sprintf("%d", j)
				}
				ans = append(ans, fmt.Sprintf("%d:%s", i, min))
			}
		}
	}
	return ans
}

func getCountOfOne(n int) int {
	count := 0
	for n > 0 {
		count += n & 1
		n >>= 1
	}
	return count
}
