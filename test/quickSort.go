package test

import (
	"fmt"
)

func Run_quickSort() {
	fmt.Println(quickSort([]int{4, 1, 6, 2, 8, 4, 1, 9}))
}

func quickSort(nums []int) []int {
	return quickSort_helper(nums, 0, len(nums)-1)
}

func quickSort_helper(nums []int, l, r int) []int {
	if l+1 >= r {
		return nums
	}

	first, last := l, r-1
	key := nums[first]
	for first < last {
		for first < last && nums[last] >= key {
			last--
		}
		nums[first] = nums[last]
		for first < last && nums[first] < key {
			first++
		}
		nums[last] = nums[first]
	}
	nums[first] = key
	nums = quickSort_helper(nums, l, first)
	nums = quickSort_helper(nums, first+1, r)
	return nums
}
