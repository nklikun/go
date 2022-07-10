package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/shu-zu-zhong-shu-zi-chu-xian-de-ci-shu-lcof/
func Run56_1() {
	fmt.Println(singleNumbers([]int{1, 4, 1, 6}))
}

func singleNumbers(nums []int) []int {
	xor := 0
	for _, n := range nums {
		xor ^= n
	}
	key := 1
	for (xor & key) == 0 {
		key = key << 1
	}
	ans := make([]int, 2)
	for _, n := range nums {
		if (n & key) != 0 {
			ans[0] ^= n
		} else {
			ans[1] ^= n
		}
	}
	return ans
}
