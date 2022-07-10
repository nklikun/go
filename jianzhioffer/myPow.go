package jianzhioffer

// https://leetcode.cn/problems/shu-zhi-de-zheng-shu-ci-fang-lcof/
func Run16() {

}

func myPow(x float64, n int) float64 {
	negative := false
	if n < 0 {
		n = -n
		negative = true
	}
	ans := 1.0
	x_con := x
	for n > 0 {
		if n%2 == 1 {
			ans *= x_con
		}
		x_con *= x_con
		n = n >> 1
	}
	if negative {
		return 1 / ans
	}
	return ans
}
