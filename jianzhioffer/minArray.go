package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/xuan-zhuan-shu-zu-de-zui-xiao-shu-zi-lcof/
func Run11() {
	fmt.Println(minArray([]int{3, 4, 5, 1, 2}))
}

func minArray(numbers []int) int {
	l, r, mid := 0, len(numbers)-1, 0
	for l < r {
		mid = (l + r) / 2
		if numbers[r] > numbers[mid] {
			r = mid
		} else if numbers[r] < numbers[mid] {
			l = mid + 1
		} else {
			r--
		}
	}
	return numbers[l]
}
