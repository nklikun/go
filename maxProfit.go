package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock
func run121() {

}

func maxProfit(prices []int) int {
	s, max := 0, 0
	for i := 0; i < len(prices); i++ {
		if prices[i] > prices[s] {
			earn := prices[i] - prices[s]
			if earn > max {
				max = earn
			}
		} else {
			s = i
		}
	}
	return max
}
