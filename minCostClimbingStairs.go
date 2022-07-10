package main

// https://leetcode.cn/problems/min-cost-climbing-stairs
func run746() {
	minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1})
}

func minCostClimbingStairs(cost []int) int {
	l := len(cost)
	sum := make([]int, l)
	sum[0] = cost[0]
	sum[1] = cost[1]
	for i := 2; i < l; i++ {
		sum[i] += min746(sum[i-1], sum[i-2])
		sum[i] += cost[i]
	}
	return min746(sum[l-1], sum[l-2])
}

func min746(a, b int) int {
	if a < b {
		return a
	}
	return b
}
