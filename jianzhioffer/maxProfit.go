package jianzhioffer

// https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/
func Run63() {}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	low, profit := prices[0], ^(int(^uint32(0) >> 1))
	for _, price := range prices {
		if price < low {
			low = price
		} else {
			tmp := price - low
			if tmp > profit {
				profit = tmp
			}
		}
	}
	return profit
}
