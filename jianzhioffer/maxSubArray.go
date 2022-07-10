package jianzhioffer

// https://leetcode.cn/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/
func Run42() {}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sum, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = 0
		}
		sum += nums[i]
		if sum > max {
			max = sum
		}
	}

	return max
}
