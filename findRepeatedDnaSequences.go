package main

// https://leetcode.cn/problems/repeated-dna-sequences/solution/
func run187() {}

func findRepeatedDnaSequences(s string) []string {
	if len(s) <= 10 {
		return nil
	}
	subs := make(map[string]int, len(s)-9)
	for i := 0; i <= len(s)-10; i++ {
		subs[s[i:i+10]]++
	}
	ans := make([]string, 0)
	for k, v := range subs {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}
