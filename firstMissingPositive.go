package main

import "fmt"

/*
41. 缺失的第一个正数
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。


示例 1：

输入：nums = [1,2,0]
输出：3
示例 2：

输入：nums = [3,4,-1,1]
输出：2
示例 3：

输入：nums = [7,8,9,11,12]
输出：1


提示：

1 <= nums.length <= 5 * 105
-231 <= nums[i] <= 231 - 1
*/
func rum41() {
	// fmt.Println(firstMissingPositive([]int{3, 4, -1, 1}))
	// fmt.Println(firstMissingPositive([]int{2, 2}))
	fmt.Println(firstMissingPositive([]int{0, 1, 2}))
}

// 优化：使用nums本身存储index，节省空间
func firstMissingPositive(nums []int) int {
	// 试着使nums[i]存储的内容为i+1
	// 依次进行交换，一旦交换不下去（index不在i~n-1）或者交换过的（已经归位），则下一个
	for i := 0; i < len(nums); i++ {
		k := nums[i] - 1
		for k >= 0 && k < len(nums) && nums[i] != i+1 && nums[i] != nums[k] {
			nums[i], nums[k] = nums[k], nums[i]
			k = nums[i] - 1
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

/*
// time: O(n), space: O(n)
func firstMissingPositive(nums []int) int {
	visited := make([]bool, len(nums)+1)
	for _, num := range nums {
		if num >= 1 && num <= len(nums) {
			visited[num] = true
		}
	}
	for i := 1; i <= len(nums); i++ {
		if !visited[i] {
			return i
		}
	}
	return len(nums) + 1
}
*/
