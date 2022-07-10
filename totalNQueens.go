package main

// https://leetcode.cn/problems/n-queens-ii/
func run52() {

}

func totalNQueens(n int) int {
	// 左下到右上的斜线
	upwards := 0
	// 左上到右下的斜线
	downwards := 0
	// 竖直方向
	vertical := 0

	total := 0
	var helper func(int, int, int, int, int)
	helper = func(up, down, ver int, rows int, n int) {
		if rows == n {
			total += 1
			return
		}

		l := rows
		for i := 0; i < n; i++ {
			if (ver&(1<<i)) == 0 && (up&(1<<(l+i))) == 0 && (down&(1<<(n-1+l-i))) == 0 {
				ver += 1 << i
				up += 1 << (l + i)
				down += 1 << (n - 1 + l - i)

				helper(up, down, ver, l+1, n)

				// backtrace, removed the added
				ver -= 1 << i
				up -= 1 << (l + i)
				down -= 1 << (n - 1 + l - i)
			}
		}
	}

	helper(upwards, downwards, vertical, 0, n)
	return total
}
