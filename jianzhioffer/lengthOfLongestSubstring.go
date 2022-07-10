package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/zui-chang-bu-han-zhong-fu-zi-fu-de-zi-zi-fu-chuan-lcof/
func Run48() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {
	l, max := 0, 0
	chars := make(map[byte]bool, 0)
	for i := 0; i < len(s); i++ {
		_, ok := chars[s[i]]
		if !ok {
			chars[s[i]] = true
			if i-l+1 > max {
				max = i - l + 1
			}
		} else {
			for ; l < i; l++ {
				if s[l] == s[i] {
					l++
					break
				}
				delete(chars, s[l])
			}
		}
	}
	return max
}
