package jianzhioffer

// https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
func Run46() {}

func translateNum(num int) int {
	if num < 10 || (num >= 26 && num < 100) {
		return 1
	}
	if num < 26 {
		return 2
	}

	total := translateNum(num / 10)
	if num%100 < 26 && num%100 > 9 {
		total += translateNum(num / 100)
	}
	return total
}

// func translateNum(num int) int {
// 	n := int32(num)
// 	if n < 10 || (n >= 26 && n < 100) {
// 		return 1
// 	}
// 	if n < 26 {
// 		return 2
// 	}

// 	total := translateNum(int(n / 10))
// 	if n%100 < 26 && n%100 > 9 {
// 		total += translateNum(int(n / 100))
// 	}
// 	return total
// }
