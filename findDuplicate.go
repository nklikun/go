package main

// https://leetcode.cn/problems/find-the-duplicate-number/
func run287() {}

func findDuplicate(nums []int) int {
	n := len(nums) - 1
	res := 0
	for i := 0; i < 32; i++ {
		cnt1, cnt2 := 0, 0
		for _, num := range nums {
			if (num>>i)&1 != 0 {
				cnt1++
			}
		}
		for j := 1; j <= n; j++ {
			if (j>>i)&1 != 0 {
				cnt2++
			}
		}
		if cnt1 > cnt2 {
			res |= (1 << i)
		}
	}
	return res
}
