package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/he-wei-sde-lian-xu-zheng-shu-xu-lie-lcof/
func Run57_2() {
	fmt.Println(findContinuousSequence(9))
}

func findContinuousSequence(target int) [][]int {
	ans := make([][]int, 0)
	limit := target / 2
	sum := 0
	l, r := 1, 1
	for l <= limit {
		if target > sum {
			sum += r
			r++
		} else if target < sum {
			sum -= l
			l++
		} else {
			seq := make([]int, r-l)
			for i := l; i < r; i++ {
				seq[i-l] = i
			}
			ans = append(ans, seq)
			sum -= l
			l++
			sum += r
			r++
		}
	}
	return ans
}
