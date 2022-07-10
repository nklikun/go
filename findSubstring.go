package main

import "fmt"

// https://leetcode.cn/problems/substring-with-concatenation-of-all-words
func run30() {
	s := "wordgoodgoodgoodbestword"
	a := findSubstring(s, []string{"word", "good", "best", "good"})
	fmt.Println(a)
}

func findSubstring(s string, words []string) []int {
	length := len(words) * len(words[0])
	indexes := make([]int, 0)

	for i := 0; i <= len(s)-length; i++ {
		//save words to map, mmap[word name]appear count
		wordsMap := make(map[string]int)
		for _, word := range words {
			if _, ok := wordsMap[word]; ok {
				wordsMap[word]++
			} else {
				wordsMap[word] = 1
			}
		}

		if isSubStringMatch(s[i:i+length], wordsMap, len(words[0]), len(words)) {
			indexes = append(indexes, i)
		}
	}

	return indexes
}

func isSubStringMatch(s string, words map[string]int, singleLen, wordsCnt int) bool {
	for i := 0; i < wordsCnt; i++ {
		w := s[i*singleLen : (i+1)*singleLen]
		if val, ok := words[w]; ok {
			if val > 1 {
				words[w] = val - 1
			} else {
				delete(words, w)
			}
		} else {
			return false
		}
	}
	return true
}
