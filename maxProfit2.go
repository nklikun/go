package main

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func run122() {

}

func maxProfit2(prices []int) int {
	l := len(prices)
	total := 0
	for i := 1; i < l; i++ {
		profit := prices[i] - prices[i-1]
		if profit > 0 {
			total += profit
		}
	}
	return total
}
