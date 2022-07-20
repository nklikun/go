package main

import "fmt"

/*
214. 最短回文串
给定一个字符串 s，你可以通过在字符串前面添加字符将其转换为回文串。
找到并返回可以用这种方式转换的最短回文串。

示例 1：

输入：s = "aacecaaa"
输出："aaacecaaa"
示例 2：

输入：s = "abcd"
输出："dcbabcd"


提示：

0 <= s.length <= 5 * 104
s 仅由小写英文字母组成
*/
func run214() {
	fmt.Println(shortestPalindrome("aacecaaa") == "aaacecaaa")
	fmt.Println(shortestPalindrome("abcd") == "dcbabcd")
	fmt.Println(shortestPalindrome("abbac") == "cabbac")
	fmt.Println(shortestPalindrome("acabacad") == "dacabacad")
	fmt.Println(shortestPalindrome("abba") == "abba")
	fmt.Println(shortestPalindrome("aba") == "aba")
	fmt.Println(shortestPalindrome("aaabba") == "abbaaabba")

	fmt.Println(shortestPalindrome("ba") == "aba")
	fmt.Println(shortestPalindrome("abb") == "bbabb")
	fmt.Println(shortestPalindrome("aabba") == "abbaabba")
}

func shortestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	mid := len(s) / 2
	if len(s)%2 == 0 {
		mid--
	}
	for i := mid; i >= 0; i-- {
		added := ""

		// try i == i+1, i-1 == i+2, ...
		j := 0
		k := i + 1 // 左：i-j, 右：i+1+j
		for ; ; j++ {
			if i-j < 0 {
				if k+j == len(s) {
					return added + s
				}
				added = string(s[k+j]) + added
			} else if k+j >= len(s) || s[i-j] != s[k+j] {
				break
			}
		}

		// try i+1 == i-1, ...
		j = 1 //skip j == 0, 即 i == i
		for ; ; j++ {
			if i-j < 0 {
				if i+j == len(s) {
					return added + s
				}
				added = string(s[i+j]) + added
			} else if i+j >= len(s) || s[i-j] != s[i+j] {
				break
			}
		}
	}
	return ""
}
