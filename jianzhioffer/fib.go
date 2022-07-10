package jianzhioffer

// https://leetcode.cn/problems/fei-bo-na-qi-shu-lie-lcof/
func Run10_1() {}

func fib(n int) int {
	if n < 2 {
		return n
	}

	mod := 1000000007
	a, b, sum := 0, 0, 1
	for i := 2; i <= n; i++ {
		a = b
		b = sum
		sum = (a + b) % mod
	}

	return sum
}
