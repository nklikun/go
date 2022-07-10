package main

// https://leetcode.cn/problems/find-the-town-judge/
func run997() {

}

func findJudge(n int, trust [][]int) int {
	trusts := make([]int, n+1)
	trusted := make([]int, n+1)
	for _, t := range trust {
		trusts[t[0]]++
		trusted[t[1]]++
	}

	for i := 1; i < n+1; i++ {
		if trusts[i] == 0 && trusted[i] == n-1 {
			return i
		}
	}
	return -1
}
