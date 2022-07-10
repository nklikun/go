package main

import "fmt"

// https://leetcode.cn/problems/next-permutation
func run31() {
	nums := []int{1, 3, 2}
	fmt.Println(nums)
	nextPermutation(nums)
	fmt.Println(nums)

	// 1.3.2->2.1.3

}

func nextPermutation(nums []int) {
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			// 找到可以交换的位置， 比如3 of 1, 2, 3, 5, 4, 2
			// 从当前数字后面，找到比当前数字大的最小的数(例中为4，不能是5或者2)，进行交换
			min := nums[i+1]
			minIndex := i + 1
			for j := i + 1; j < len(nums); j++ {
				if min > nums[j] && nums[j] > nums[i] {
					min = nums[j]
					minIndex = j
				}
			}
			tmp := nums[i]
			nums[i] = nums[minIndex]
			nums[minIndex] = tmp

			// 交换完成后把j+1往后部分进行升序排序
			for m := i + 1; m < len(nums)-1; m++ {
				for n := m + 1; n < len(nums); n++ {
					if nums[m] > nums[n] {
						t := nums[m]
						nums[m] = nums[n]
						nums[n] = t
					}
				}
			}

			return
		}
	}

	// 没有发现可换的数字，进行倒序
	i := 0
	j := len(nums) - 1
	for i < j {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
		i++
		j--
	}
}
