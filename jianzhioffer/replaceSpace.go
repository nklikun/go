package jianzhioffer

import "bytes"

// https://leetcode.cn/problems/ti-huan-kong-ge-lcof/
func Run05() {}

func replaceSpace(s string) string {
	b := new(bytes.Buffer)

	for _, c := range s {
		if c == ' ' {
			b.WriteByte('%')
			b.WriteByte('2')
			b.WriteByte('0')
		} else {
			b.WriteByte(byte(c))
		}
	}
	return b.String()
}
