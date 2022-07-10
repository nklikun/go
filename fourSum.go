package main

import (
	"fmt"
	"sort"
)

// https://leetcode.cn/problems/4sum
func run18() {
	nums := []int{1, 0, -1, 0, -2, 2}
	fmt.Println(fourSum(nums, 0))
}

func fourSum(nums []int, target int) [][]int {
	rets := make([][]int, 0)
	n := len(nums)
	if n < 4 {
		return rets
	}

	dic := make(map[string]int)

	sort.Ints(nums)
	for i := 0; i < n-3; i++ {
		for j := i + 1; j < n-2; j++ {
			l := j + 1
			r := n - 1

			for r > l {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					ret := []int{nums[i], nums[j], nums[l], nums[r]}
					key := fmt.Sprintf("%d%d%d%d", nums[i], nums[j], nums[l], nums[r])
					if _, exist := dic[key]; !exist {
						rets = append(rets, ret)
						dic[string(key)] = 0
					}
					// if number 1 and number 2 are kept, number3 and number4 have to change
					l++
					r--
				} else if sum > target {
					r--
				} else {
					l++
				}

			}

		}
	}

	return rets
}
