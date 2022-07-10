package niuke

import (
	"fmt"
	"sort"
)

// https://www.nowcoder.com/practice/a43a2b986ef34843ac4fdd9159b69863?tpId=295&tqId=700&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM56() {
	fmt.Println(permuteUnique([]int{1, 1, 2}))
}

/**
 *
 * @param num int整型一维数组
 * @return int整型二维数组
 */
func permuteUnique(num []int) [][]int {
	visited := make([]bool, len(num))
	// write code here
	ans := make([][]int, 0)
	array := make([]int, 0)
	var dfs func()
	dfs = func() {
		if len(array) == len(num) {
			list := make([]int, len(array))
			copy(list, array)
			ans = append(ans, list)
			return
		}
		for i := 0; i < len(num); i++ {
			if visited[i] || (i > 0 && num[i] == num[i-1] && !visited[i-1]) {
				continue
			}
			array = append(array, num[i])
			visited[i] = true
			dfs()
			visited[i] = false
			array = array[:len(array)-1]
		}
	}
	sort.Ints(num)
	dfs()
	return ans
}
