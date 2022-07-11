package main

import "fmt"

// https://leetcode.cn/problems/single-number-ii/
func run137() {
	// fmt.Println(singleNumber([]int{0, 1, 0, 1, 0, 1, 99}))
	fmt.Println(singleNumber([]int{-2, -2, 1, 1, 4, 1, 4, 4, -4, -2}))
}

func singleNumber(nums []int) int {
	n := int32(0)
	for i := 0; i < 32; i++ {
		count := int32(0)
		for _, num := range nums {
			count += (int32(num) >> i) & 1
		}
		if count%3 > 0 {
			n |= 1 << i
		}
	}
	return int(n)
}

// https://leetcode.cn/problems/single-number-iii/submissions/
func run260() {}
func singleNumber3(nums []int) []int {
	twoSum := 0
	for _, num := range nums {
		twoSum ^= num
	}
	pos := 0
	for ; (twoSum>>pos)&1 == 0; pos++ {
	}

	num1, num2 := 0, 0
	for _, num := range nums {
		if (num>>pos)&1 == 1 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
