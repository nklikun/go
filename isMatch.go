package main

import (
	"fmt"
	"time"
)

// https://leetcode.cn/problems/wildcard-matching
func run44() {
	run44_help("aa", "a")
	run44_help("aa", "*")
	run44_help("cb", "*a")
	run44_help("adceb", "*a*b")
	run44_help("acdcb", "a*c?b")
	/*
		"babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab"
		"***bba**a*bbba**aab**b"
	*/
	run44_help("babaaababaabababbbbbbaabaabbabababbaababbaaabbbaaab", "***bba**a*bbba**aab**b")
}

func run44_help(s, p string) {
	start := time.Now()
	fmt.Println("s: ", s)
	fmt.Println("p: ", p)
	fmt.Println("match: ", isMatch(s, p))
	end := time.Now()
	fmt.Println("times: ", end.Sub(start))
}

/*func isMatch(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	if len(s) == 0 {
		return p[0] == '*' && isMatch(s, p[1:])
	}

	c := p[0]
	if c == '?' {
		return isMatch(s[1:], p[1:])
	} else if c == '*' {
		return isMatch(s[1:], p) || isMatch(s, p[1:]) || isMatch(s[1:], p[1:])
	} else {
		if s[0] != p[0] {
			return false
		}
		return isMatch(s[1:], p[1:])
	}
}*/

func isMatch(s string, p string) bool {
	dps := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dps[i] = make([]int, len(p))
	}

	var dfs func(string, string, int, int) bool
	dfs = func(s string, p string, si int, pi int) bool {
		if pi == len(p) {
			return len(s) == si
		}

		if len(s) == si {
			return p[pi] == '*' && dfs(s, p, si, pi+1)
		}

		if dps[si][pi] != 0 {
			return dps[si][pi] == 1
		}

		c := p[pi]
		if c == '?' {
			if dfs(s, p, si+1, pi+1) {
				dps[si][pi] = 1
			} else {
				dps[si][pi] = -1
			}
		} else if c == '*' {
			if dfs(s, p, si+1, pi) || dfs(s, p, si, pi+1) || dfs(s, p, si+1, pi+1) {
				dps[si][pi] = 1
			} else {
				dps[si][pi] = -1
			}
		} else {
			if s[si] != p[pi] {
				dps[si][pi] = -1
			} else {
				if dfs(s, p, si+1, pi+1) {
					dps[si][pi] = 1
				} else {
					dps[si][pi] = -1
				}
			}
		}
		return dps[si][pi] == 1
	}
	return dfs(s, p, 0, 0)
}
