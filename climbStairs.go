package main

// https://leetcode.cn/problems/climbing-stairs
func run70() {

}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}

	dps := make([]int, n+1)
	dps[0] = 1
	dps[1] = 1
	for i := 2; i <= n; i++ {
		dps[i] = dps[i-1] + dps[i-2]
	}

	return dps[n]
}
