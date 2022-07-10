package main

// https://leetcode.cn/problems/maximum-subarray
func run53() {

}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	sum := 0
	for _, num := range nums {
		sum += num
		if sum > max {
			max = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return max
}
