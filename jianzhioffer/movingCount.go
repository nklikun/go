package jianzhioffer

// https://leetcode.cn/problems/ji-qi-ren-de-yun-dong-fan-wei-lcof/
func Run13() {

}

func movingCount(m int, n int, k int) int {
	passed := make([][]bool, m)
	for i := 0; i < m; i++ {
		passed[i] = make([]bool, n)
	}
	var dfs func(int, int)
	dfs = func(i int, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || passed[i][j] {
			return
		}
		sum := i%10 + i/10 + j%10 + j/10
		if sum > k {
			return
		}
		passed[i][j] = true
		dfs(i-1, j)
		dfs(i, j-1)
		dfs(i+1, j)
		dfs(i, j+1)
	}
	dfs(0, 0)
	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if passed[i][j] {
				total++
			}
		}
	}
	return total
}
