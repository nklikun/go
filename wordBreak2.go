package main

import "fmt"

// https://leetcode.cn/problems/word-break-ii/
func run140() {
	ret1 := wordBreak2("catsanddog", []string{"cat", "cats", "and", "sand", "dog"})
	for _, str := range ret1 {
		fmt.Println(str)
	}

	ret2 := wordBreak2("pineapplepenapple", []string{"apple", "pen", "applepen", "pine", "pineapple"})
	for _, str := range ret2 {
		fmt.Println(str)
	}
}

func wordBreak2(s string, wordDict []string) []string {
	if len(wordDict) == 0 || len(s) == 0 {
		return nil
	}

	dic := make(map[string]bool)
	minLen := len(wordDict[0])
	maxLen := minLen
	for _, w := range wordDict {
		dic[w] = true
		if len(w) < minLen {
			minLen = len(w)
		}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	var wordBreak2_helper func(string) []string
	wordBreak2_helper = func(str string) []string {
		ls := len(str)
		sIndex, eIndex := ls-maxLen, ls-minLen
		if sIndex < 0 {
			sIndex = 0
		}
		if eIndex < 0 {
			return nil
		}

		ans := make([]string, 0)
		for i := sIndex; i <= eIndex; i++ {
			// fmt.Println(str[i:len(str)])
			if dic[str[i:len(str)]] {
				if i == 0 {
					ans = append(ans, str[i:ls])
					continue
				}
				befores := wordBreak2_helper(str[0:i])
				for _, b := range befores {
					ans = append(ans, fmt.Sprintf("%s %s", b, str[i:ls]))
				}
			}
		}
		return ans
	}
	return wordBreak2_helper(s)
}

/*
func wordBreak(s string, wordDict []string) []string {
	if len(wordDict) == 0 || len(s) == 0 {
		return nil
	}

	dic := make(map[string]bool)
	minLen := len(wordDict[0])
	maxLen := minLen
	for _, w := range wordDict {
		dic[w] = true
		if len(w) < minLen {
			minLen = len(w)
		}
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	var wordBreak_helper func(string) []string
	wordBreak_helper = func(str string) []string {
		ls := len(str)
		sIndex, eIndex := ls-maxLen, ls-minLen
		if sIndex < 0 {
			sIndex = 0
		}
		if eIndex < 0 {
			return nil
		}

		ans := make([]string, 0)
		for i := sIndex; i <= eIndex; i++ {
			// fmt.Println(str[i:len(str)])
			if dic[str[i:len(str)]] {
				if i == 0 {
					ans = append(ans, str[i:ls])
					continue
				}
				befores := wordBreak_helper(str[0:i])
				for _, b := range befores {
					ans = append(ans, fmt.Sprintf("%s %s", b, str[i:ls]))
				}
			}
		}
		return ans
	}
	return wordBreak_helper(s)
}
*/
