package main

import (
	"fmt"
	"time"
)

// https://leetcode.cn/problems/n-queens/
func run51() {
	start := time.Now()
	fmt.Println(solveNQueens(12))
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	// 左下到右上的斜线
	upwards := 0
	// 左上到右下的斜线
	downwards := 0
	// 竖直方向
	vertical := 0

	var helper func(int, int, int, []int, int)
	helper = func(up, down, ver int, filled []int, n int) {
		l := len(filled)
		if l == n {
			matrix := make([]string, n)
			for i := 0; i < n; i++ {
				row := ""
				pos := filled[i]
				for j := 0; j < n; j++ {
					if j == pos {
						row += "Q"
					} else {
						row += "."
					}
				}
				matrix[i] = row
			}
			ans = append(ans, matrix)
			return
		}

		for i := 0; i < n; i++ {
			if (ver&(1<<i)) == 0 && (up&(1<<(l+i))) == 0 && (down&(1<<(n-1+l-i))) == 0 {
				ver += 1 << i
				up += 1 << (l + i)
				down += 1 << (n - 1 + l - i)

				// filled不需要deepcopy，因为决定其内容的是ver,up,down,后续更改的内容会被覆盖
				filled = append(filled, i)
				helper(up, down, ver, filled, n)

				// backtrace, removed the added
				ver -= 1 << i
				up -= 1 << (l + i)
				down -= 1 << (n - 1 + l - i)
				filled = filled[:len(filled)-1]
			}
		}
	}

	helper(upwards, downwards, vertical, []int{}, n)
	return ans
}

/*
func solveNQueens(n int) [][]string {
	ans := make([][]string, 0)
	// 左下到右上的斜线
	upwards := make([]bool, n*2-1)
	// 左上到右下的斜线
	downwards := make([]bool, n*2-1)
	// 竖直方向
	vertical := make([]bool, n)
	line := ""
	for i := 0; i < n; i++ {
		line += "."
	}

	var helper func([]bool, []bool, []bool, []string, int)
	helper = func(up, down, ver []bool, filled []string, n int) {
		l := len(filled)
		if l == n {
			ans = append(ans, filled)
			return
		}

		for i := 0; i < n; i++ {
			v := solveNQueens_deepCopyBool(ver)
			d := solveNQueens_deepCopyBool(down)
			u := solveNQueens_deepCopyBool(up)
			f := solveNQueens_deepCopyString(filled)
			if !v[i] && !up[l+i] && !d[n-1+l-i] {
				v[i] = true
				u[l+i] = true
				fmt.Println("upwards: ", l+i)
				d[n-1+l-i] = true
				fmt.Println("downwards: ", n-1+l-i)

				newline := line
				if i == 0 {
					newline = "Q" + newline[1:]
				} else if i == n-1 {
					newline = newline[:n-1] + "Q"
				} else {
					newline = newline[0:i] + "Q" + newline[i+1:n]
				}
				f := append(f, newline)
				helper(u, d, v, f, n)
			}
		}
	}

	helper(upwards, downwards, vertical, []string{}, n)
	return ans
}

func solveNQueens_deepCopyBool(src []bool) []bool {
	dst := make([]bool, len(src))
	for i, v := range src {
		dst[i] = v
	}
	return dst
}

func solveNQueens_deepCopyString(src []string) []string {
	dst := make([]string, len(src))
	for i, v := range src {
		dst[i] = v
	}
	return dst
}
*/
