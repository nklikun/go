package jianzhioffer

// https://leetcode.cn/problems/qing-wa-tiao-tai-jie-wen-ti-lcof/
func Run10_2() {}

func numWays(n int) int {
	if n < 2 {
		return 1
	}

	mod := 1000000007
	p, q, tmp := 1, 1, 0
	for i := 2; i <= n; i++ {
		tmp = q
		q = (p + q) % mod
		p = tmp
	}

	return q
}
