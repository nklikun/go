package jianzhioffer

// https://leetcode.cn/problems/ju-zhen-zhong-de-lu-jing-lcof/
func Run12() {

}

func exist(board [][]byte, word string) bool {
	m, n, k := len(board), len(board[0]), len(word)
	used := make([][]bool, m)
	for i := 0; i < m; i++ {
		used[i] = make([]bool, n)
	}

	var dfs func(int, int, int) bool
	dfs = func(i int, j int, pos int) bool {
		// i: board[i], j: board[j], pos: word index, return: exist/not exist
		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}
		if board[i][j] != word[pos] || used[i][j] {
			return false
		}
		if pos == k-1 {
			return true
		}

		used[i][j] = true
		if dfs(i, j-1, pos+1) || dfs(i-1, j, pos+1) || dfs(i+1, j, pos+1) || dfs(i, j+1, pos+1) {
			return true
		}
		used[i][j] = false
		return false
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
