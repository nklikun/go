package main

import (
	"fmt"
	"sort"
)

/*
给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的每个数字在每个组合中只能使用 一次 。

注意：解集不能包含重复的组合。



示例 1:

输入: candidates = [10,1,2,7,6,1,5], target = 8,
输出:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]
示例 2:

输入: candidates = [2,5,2,1,2], target = 5,
输出:
[
[1,2,2],
[5]
]


提示:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/combination-sum-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func run40() {
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	ans := make([][]int, 0)
	cur := make([]int, 0)
	used := make([]bool, len(candidates))
	var dfs func(int, int)
	dfs = func(k int, tar int) {
		if tar == 0 {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			ans = append(ans, tmp)
			return
		}

		// candicateds is sorted, so if candidates[k] > tar, others left are larger too
		if k == len(candidates) || candidates[k] > tar {
			// no num left
			return
		}

		// append k
		// skip duplicated, or tar is not enough for add num k
		if !(k > 0 && candidates[k] == candidates[k-1] && !used[k-1]) {
			used[k] = true
			cur = append(cur, candidates[k])
			dfs(k+1, tar-candidates[k])
			used[k] = false
			cur = cur[:len(cur)-1]
		}

		// not append k
		dfs(k+1, tar)
	}
	dfs(0, target)
	return ans
}
