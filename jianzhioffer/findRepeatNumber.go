package jianzhioffer

// https://leetcode.cn/problems/shu-zu-zhong-zhong-fu-de-shu-zi-lcof/
func Run03() {}

func findRepeatNumber(nums []int) int {
	numMap := make(map[int]bool, 0)
	for _, num := range nums {
		if _, ok := numMap[num]; ok {
			return num
		}
		numMap[num] = true
	}
	return -1
}
