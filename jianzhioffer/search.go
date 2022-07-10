package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/zai-pai-xu-shu-zu-zhong-cha-zhao-shu-zi-lcof/
func Run53_1() {
	fmt.Println(search([]int{1, 2, 2, 2, 2, 2, 3}, 2))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l] == target && nums[r] == target {
			return r - l + 1
		}
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			if nums[l] != target {
				l = mid
				for ; nums[l] == target; l-- {
				}
				l++
			} else {
				r = mid
				for ; nums[r] == target; r++ {
				}
				r--
			}
		}
	}
	return r - l + 1
}
