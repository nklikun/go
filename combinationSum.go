package main

import "fmt"

// https://leetcode.cn/problems/combination-sum/solution/
func run39() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}

	ans := make([][]int, 0)
	for i := 0; candidates[0]*i <= target; i++ {
		if candidates[0]*i == target {
			r := make([]int, 0)
			for j := 0; j < i; j++ {
				r = append(r, candidates[0])
			}
			ans = append(ans, r)
			break
		}
		res := combinationSum(candidates[1:], target-i*candidates[0])
		for _, r := range res {
			for j := 0; j < i; j++ {
				r = append([]int{candidates[0]}, r...)
			}
			ans = append(ans, r)
		}
	}
	return ans
}
