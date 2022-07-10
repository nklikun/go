package niuke

import (
	"fmt"
	"sort"
	"time"
)

// https://www.nowcoder.com/practice/fe6b651b66ae47d7acce78ffdd9a96c7?tpId=295&tqId=23291&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM58() {
	start := time.Now()
	fmt.Println(len(Permutation("qwertyuio")))
	end := time.Now()
	fmt.Println("times: ", end.Sub(start))
	fmt.Println(Permutation("aab"))
}

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param str string字符串
 * @return string字符串一维数组
 */
func Permutation(str string) []string {
	// write code here
	if len(str) == 0 {
		return nil
	}
	if len(str) == 1 {
		return []string{str}
	}

	s := []byte(str)
	sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
	str = string(s)
	ans := make([]string, 0)
	for i := 0; i < len(str); i++ {
		if i > 0 && str[i] == str[i-1] {
			continue
		}
		rets := Permutation(str[:i] + str[i+1:])
		for j := 0; j < len(rets); j++ {
			ans = append(ans, string(str[i])+rets[j])
		}
	}
	return ans
}
