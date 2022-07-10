package main

import (
	"sort"
)

// https://leetcode.cn/problems/group-anagrams/
func Run49() {}

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string, 0)
	for _, str := range strs {
		sorted := []byte(str)
		sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })
		key := string(sorted)
		strMap[key] = append(strMap[key], str)
	}
	ans := make([][]string, 0, len(strMap))
	for _, v := range strMap {
		ans = append(ans, v)
	}
	return ans
}
