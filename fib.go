package main

// https://leetcode.cn/problems/fibonacci-number
func run509() {

}

func fib(n int) int {
	fn := make([]int, n+1)
	fn[0] = 0
	fn[1] = 1
	for i := 2; i <= n; i++ {
		fn[i] = fn[i-1] + fn[i-2]
	}
	return fn[n]
}
