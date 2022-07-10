package niuke

import (
	"fmt"
)

// https://www.nowcoder.com/practice/3911a20b3f8743058214ceaa099eeb45?tpId=295&tqId=988994&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM70() {
	fmt.Print(minMoney([]int{5, 2, 3}, 20))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 最少货币数
 * @param arr int整型一维数组 the array
 * @param aim int整型 the target
 * @return int整型
 */
func minMoney(arr []int, aim int) int {
	// write code here
	if len(arr) == 0 {
		return -1
	}
	if aim <= 0 {
		return 0
	}

	counts := make([]int, aim+1)
	for i := 1; i < aim+1; i++ {
		counts[i] = aim + 1
	}

	for i := 1; i <= aim; i++ {
		for j := 0; j < len(arr); j++ {
			if arr[j] <= i {
				if counts[i] > counts[i-arr[j]]+1 {
					counts[i] = counts[i-arr[j]] + 1
				}
			}
		}
	}
	if counts[aim] > aim {
		return -1
	}
	return counts[aim]
}
