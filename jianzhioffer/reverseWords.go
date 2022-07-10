package jianzhioffer

import (
	"fmt"
)

// https://leetcode.cn/problems/fan-zhuan-dan-ci-shun-xu-lcof/
func Run58_1() {
	fmt.Println(reverseWords("      the sky    is   blue  "))
}

func reverseWords(s string) string {
	l, r := 0, len(s)-1
	for s[l] == ' ' {
		l++
	}
	for s[r] == ' ' {
		r--
	}
	res, tmp := "", ""
	for ; l <= r; l++ {
		if s[l] == ' ' && len(tmp) > 0 {
			res = " " + tmp + res
			tmp = ""
			continue
		}
		if s[l] != ' ' {
			tmp += string(s[l])
		}
	}
	return tmp + res
}

// func reverseWords(s string) string {
// 	spaceRelace, _ := regexp.Compile(`\s+`)
// 	strs := spaceRelace.Split(s, -1)
// 	res := ""
// 	for i := len(strs) - 1; i >= 0; i-- {
// 		if strs[i] != "" {
// 			res += " "
// 			res += strs[i]
// 		}
// 	}
// 	if len(res) == 0 {
// 		return ""
// 	}
// 	return res[1:]
// }
