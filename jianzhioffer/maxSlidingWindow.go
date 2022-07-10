package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/hua-dong-chuang-kou-de-zui-da-zhi-lcof/
func Run59_1() {
	fmt.Println(maxSlidingWindow([]int{1, 2, 3, 4, 5}, 2))
}

// func maxSlidingWindow(nums []int, k int) []int {
// 	if len(nums) == 0 || k > len(nums) {
// 		return nil
// 	}

// 	q := &MaxQueue{nums: make([]int, 0), ordered: make([]int, 0)}
// 	for i := 0; i < k; i++ {
// 		q.Push_back(nums[i])
// 	}

// 	ans := make([]int, len(nums)-k+1)
// 	ans[0] = q.Max_value()
// 	for i := k; i < len(nums); i++ {
// 		q.Pop_front()
// 		q.Push_back(nums[i])
// 		ans[i-k+1] = q.Max_value()
// 	}
// 	return ans
// }

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k > len(nums) {
		return nil
	}

	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[i]
		for j := i - 1; j > i-k && j >= 0; j-- {
			if ans[j] < nums[i] {
				ans[j] = nums[i]
			} else {
				break
			}
		}
	}
	return ans[:len(nums)-k+1]
}
