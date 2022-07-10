package jianzhioffer

import (
	"strconv"
	"strings"
)

// https://leetcode.cn/problems/ba-shu-zu-pai-cheng-zui-xiao-de-shu-lcof/
func Run45() {

}

func minNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i, num := range nums {
		strs[i] = strconv.Itoa(num)
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i; j < len(nums); j++ {
			if minNumber_compare(strs[i], strs[j]) {
				tmp := strs[j]
				strs[j] = strs[i]
				strs[i] = tmp
			}
		}
	}
	res := strings.Join(strs, "")
	return res
}

func minNumber_compare(n1, n2 string) bool {
	sumn1n2, _ := strconv.Atoi(n1 + n2)
	sumn2n1, _ := strconv.Atoi(n2 + n1)
	if sumn1n2 > sumn2n1 {
		return true
	}
	return false
}
