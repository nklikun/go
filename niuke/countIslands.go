package niuke

import "fmt"

// https://www.nowcoder.com/practice/0c9664d1554e466aa107d899418e814e?tpId=295&tqId=1024684&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BMP57() {
	fmt.Println(islandsCount([][]byte{{'1', '1'}, {'1', '1'}}))
}

/**
 * 判断岛屿数量
 * @param grid char字符型二维数组
 * @return int整型
 */
func islandsCount(grid [][]byte) int {
	// write code here
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m := len(grid)
	n := len(grid[0])
	total := 0
	// return: 当前岛屿是否为岛屿
	// 连接到其他岛屿不算作岛屿
	var dfs func(int, int)
	dfs = func(i, j int) {
		grid[i][j] = '0'
		if i-1 >= 0 && grid[i-1][j] == '1' {
			dfs(i-1, j)
		}
		if i+1 < m && grid[i+1][j] == '1' {
			dfs(i+1, j)
		}
		if j-1 >= 0 && grid[i][j-1] == '1' {
			dfs(i, j-1)
		}
		if j+1 < n && grid[i][j+1] == '1' {
			dfs(i, j+1)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				total++
				dfs(i, j)
			}
		}
	}
	return total
}
