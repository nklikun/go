package main

/*
给你两个单词 word1 和 word2， 请返回将 word1 转换成 word2 所使用的最少操作数  。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符


示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')


提示：

0 <= word1.length, word2.length <= 500
word1 和 word2 由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/edit-distance
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run72() {}

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m*n == 0 {
		return m + n
	}
	// dp[i][j] 代表 word1 到 i 位置转换成 word2 到 j 位置需要最少步数
	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
		for j := 0; j <= n; j++ {
			if i == 0 {
				dp[0][j] = j
			} else if j == 0 {
				dp[i][0] = i
			} else {
				if word1[i-1] == word2[j-1] {
					dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
				} else {
					dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
				}
			}
		}
	}
	return dp[m][n]
}

func min(num ...int) int {
	a := num[0]
	for _, n := range num {
		if a > n {
			a = n
		}
	}
	return a
}
