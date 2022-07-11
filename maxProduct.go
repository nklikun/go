package main

import "fmt"

// https://leetcode.cn/problems/maximum-product-of-word-lengths/
func run318() {
	fmt.Println(maxProduct([]string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"}))
}

func maxProduct(words []string) int {
	hash := make([]int, len(words))
	for i, w := range words {
		for j := 0; j < len(w); j++ {
			// if (hash[i] & (1 << (w[j] - 'a'))) == 0 {
			hash[i] |= 1 << (w[j] - 'a')
			// }
		}
	}
	max := 0
	for i := 0; i < len(hash)-1; i++ {
		for j := i + 1; j < len(hash); j++ {
			if hash[i]&hash[j] == 0 && len(words[i])*len(words[j]) > max {
				max = len(words[i]) * len(words[j])
			}
		}
	}
	return max
}

// func maxProduct(words []string) int {
// 	hash := make(map[string]int)
// 	for _, w := range words {
// 		for i := 0; i < len(w); i++ {
// 			if (hash[w] & (1 << (w[i] - 'a'))) == 0 {
// 				hash[w] |= 1 << (w[i] - 'a')
// 			}
// 		}
// 	}
// 	max := 0
// 	for i := 0; i < len(words)-1; i++ {
// 		for j := i + 1; j < len(words); j++ {
// 			a, b := words[i], words[j]
// 			if len(a)*len(b) > max && (hash[a]&hash[b]) == 0 {
// 				max = len(a) * len(b)
// 			}
// 		}
// 	}
// 	return max
// }
