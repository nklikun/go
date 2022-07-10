package main

import "fmt"

// https://leetcode.cn/problems/add-binary/
func run67() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	shift := false
	res := ""
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || shift; {
		k := 0
		if shift {
			k++
		}
		shift = false
		if i >= 0 && a[i] == '1' {
			k++
		}
		if j >= 0 && b[j] == '1' {
			k++
		}
		if k == 3 {
			res = "1" + res
			shift = true
		} else if k == 2 {
			res = "0" + res
			shift = true
		} else if k == 1 {
			res = "1" + res
		} else {
			res = "0" + res
		}
		i--
		j--
	}

	return res
}
