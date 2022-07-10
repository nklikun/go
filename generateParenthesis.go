package main

import "fmt"

// https://leetcode.cn/problems/generate-parentheses
func run22() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	parenthesis := make([]string, 0)
	generateParenthesis_helper(&parenthesis, "", 0, 0, n)
	return parenthesis
}

func generateParenthesis_helper(parenthesis *[]string, str string, l, r, n int) {
	if l > n || r > n || r > l {
		return
	}

	if l == n && r == n {
		*parenthesis = append(*parenthesis, str)
		return
	}
	generateParenthesis_helper(parenthesis, str+"(", l+1, r, n)
	generateParenthesis_helper(parenthesis, str+")", l, r+1, n)
	return
}
