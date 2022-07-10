package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/biao-shi-shu-zhi-de-zi-fu-chuan-lcof/
func Run20() {
	fmt.Println("is a number: ")
	fmt.Println(isNumber("+100"))
	fmt.Println(isNumber("5e2"))
	fmt.Println(isNumber("-123"))
	fmt.Println(isNumber("3.1416"))
	fmt.Println(isNumber("-1E-16"))
	fmt.Println(isNumber("-1E+16"))
	fmt.Println(isNumber("0123"))
	fmt.Println(isNumber("123.11E12"))
	fmt.Println(isNumber("   123.11E12   "))
	fmt.Println(isNumber("0"))
	fmt.Println(isNumber("    .1  "))
	fmt.Println(isNumber("+.5"))
	fmt.Println(isNumber("1 "))
	fmt.Println(isNumber(".1"))
	fmt.Println(isNumber(" 0"))
	fmt.Println(isNumber(" .9"))
	fmt.Println(isNumber("+.8"))
	fmt.Println(isNumber("3.e3"))
	fmt.Println(isNumber("46.e3"))

	fmt.Println(isNumber("3."))

	// ---------------------------
	fmt.Println("not a number: ")
	fmt.Println(isNumber("12e"))
	fmt.Println(isNumber("1a3.14"))
	fmt.Println(isNumber("1.2.3"))
	fmt.Println(isNumber("+-5"))
	fmt.Println(isNumber("12e+5.4"))
	fmt.Println("--")
	fmt.Println(isNumber("e"))
	fmt.Println(isNumber("."))
	fmt.Println(isNumber("e9"))
	fmt.Println(isNumber(" "))
	fmt.Println(isNumber("."))
	fmt.Println("--")
	fmt.Println(isNumber("1a3.14"))
	fmt.Println(isNumber("1.2.3"))
	fmt.Println(isNumber("+-5"))
	fmt.Println(isNumber("12e+5.4"))
	fmt.Println(isNumber(".1."))
	fmt.Println("--")
	fmt.Println(isNumber("1e."))
	fmt.Println(isNumber("-."))
	fmt.Println(isNumber(".."))
	fmt.Println(isNumber(".+"))
	fmt.Println(isNumber(".e"))
	fmt.Println(isNumber("92e1740e91"))
	fmt.Println(isNumber("41234E10."))
}

func isNumber(s string) bool {
	start, end := 0, len(s)-1
	// trim head and tail
	for ; start <= end && s[start] == ' '; start++ {
	}
	for ; end >= start && s[end] == ' '; end-- {
	}

	if start > end {
		return false
	}

	isNum, isDot, isExp, isSign := false, false, false, false
	// start to judge
	for ; start <= end; start++ {
		c := s[start]
		if c == '+' || c == '-' {
			if isNum || isDot || isSign {
				return false
			}
			isSign = true
		} else if c == 'e' || c == 'E' {
			if !isNum || isExp {
				return false
			}
			isExp = true
			isDot = false
			isSign = false
			isNum = false
		} else if c == '.' {
			if isDot || isExp {
				return false
			}
			isDot = true
		} else if c >= '0' && c <= '9' {
			isNum = true
		} else {
			return false
		}
	}

	start--
	return (s[start] <= '9' && s[start] >= '0') || (isNum && s[start] == '.')
}
