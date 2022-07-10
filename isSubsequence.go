package main

import "fmt"

// https://leetcode.cn/problems/is-subsequence
func run392() {
	fmt.Println(isSubsequence("bde", "abcdef"))
}

func isSubsequence(s string, t string) bool {
	ls := len(s)
	lt := len(t)
	i, j := 0, 0
	for i < ls && j < lt {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == ls
}

// func isSubsequence(s string, t string) bool {
// 	ls := len(s)
// 	lt := len(t)
// 	if ls < lt {
// 		return false
// 	}

// 	base := uint64(10000007)
// 	shash := uint64(0)
// 	thash := uint64(0)

// 	for i := 0; i < ls; i++ {
// 		shash = shash*base + uint64(s[i])
// 	}

// 	for i := 0; i < lt; i++ {
// 		thash = thash*base + uint64(t[i])
// 	}

// 	tbase := uint64(1)
// 	for i := 0; i < lt; i++ {
// 		tbase *= base
// 	}

// 	for i := 0; i < ls-lt; i++ {
// 		if shash == thash {
// 			return true
// 		}
// 		thash = thash*base + uint64(s[i+lt]) - uint64(s[i])*tbase
// 	}
// 	return shash == thash
// }
