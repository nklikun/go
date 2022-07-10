package main

import "fmt"

//17.电话号码的字母组合
// https://leetcode.cn/problems/letter-combinations-of-a-phone-number
func run17() {
	ret := letterCombinations("23456789")
	fmt.Println(ret)
}

func letterCombinations(digits string) []string {
	mapping := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ret := make([]string, 0)
	if len(digits) == 0 {
		return ret
	}

	// transfer '2' to 0, '3' to 1...
	index := int(digits[0]) - 50
	for _, c := range (mapping[index])[:] {
		if len(digits) == 1 {
			ret = append(ret, string(c))
		} else {
			callbackStrings := letterCombinations(digits[1:])
			for _, str := range callbackStrings {
				ret = append(ret, string(c)+str)
			}
		}
	}
	return ret
}
