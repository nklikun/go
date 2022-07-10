package niuke

import (
	"fmt"
	"sort"
)

// https://www.nowcoder.com/practice/4edf6e6d01554870a12f218c94e8a299?tpId=295&tqId=1267319&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM96() {
	// fmt.Println(minmumNumberOfHost(2, [][]int{{1, 2}, {2, 3}}))
	se := [][]int{{2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}, {2147483646, 2147483647}, {-2147483648, -2147483647}}
	fmt.Println(minmumNumberOfHost(10, se))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * 计算成功举办活动需要多少名主持人
 * @param n int整型 有n个活动
 * @param startEnd int整型二维数组 startEnd[i][0]用于表示第i个活动的开始时间，startEnd[i][1]表示第i个活动的结束时间
 * @return int整型
 */
func minmumNumberOfHost(n int, startEnd [][]int) int {
	// write code here
	times := make([]int, n*2)
	starts, ends := make([]int, n), make([]int, n)
	for i, se := range startEnd {
		// 单数存储start，双数存储end
		times[i] = se[0]
		times[n+i] = se[1]
		starts[i] = se[0]
		ends[i] = se[1]
	}
	// 排序
	sort.Ints(times)
	sort.Ints(starts)
	sort.Ints(ends)
	// 计算每一时刻主持人总数
	total, max := 0, 0
	s, e := 0, 0
	for i := 0; i < len(times); i++ {
		if i > 0 && times[i] == times[i]-1 {
			continue
		}
		for ; s < n && starts[s] == times[i]; s++ {
			total++
		}
		for ; e < n && ends[e] == times[i]; e++ {
			total--
		}
		if total > max {
			max = total
		}
	}
	return max
}
