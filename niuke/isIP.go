package niuke

import (
	"fmt"
	"strconv"
	"strings"
)

// https://www.nowcoder.com/practice/55fb3c68d08d46119f76ae2df7566880?tpId=295&tqId=1024725&ru=/exam/oj&qru=/ta/format-top101/question-ranking&sourceUrl=%2Fexam%2Foj
func Run_BM85() {
	fmt.Println(isIP("2001:0db8:85a3:0:0:8A2E:0370:7334"))
}

/**
 * 验证IP地址
 * @param IP string字符串 一个IP地址字符串
 * @return string字符串
 */
func isIP(IP string) string {
	// write code here
	array := strings.Split(IP, ".")
	if len(array) == 4 {
		for i, arr := range array {
			if !isIPV4Member(arr) {
				break
			}
			if i == 3 {
				return "IPv4"
			}
		}
	}

	array = strings.Split(IP, ":")
	if len(array) == 8 {
		for i, arr := range array {
			if !isIPV6Member(arr) {
				break
			}
			if i == 7 {
				return "IPv6"
			}
		}
	}

	return "Neither"
}

func isIPV4Member(str string) bool {
	if len(str) == 0 || len(str) > 3 {
		return false
	}
	if str[0] == '0' {
		if len(str) == 1 {
			return true
		} else {
			return false
		}
	}
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	if num > 255 || num < 0 {
		return false
	}
	return true
}

func isIPV6Member(str string) bool {
	if len(str) == 0 || len(str) > 4 {
		return false
	}
	if len(str) > 1 && str[0] == '0' && str[1] == '0' {
		return false
	}
	for i := 0; i < len(str); i++ {
		a := str[i]
		b := (a <= '9' && a >= '0') || (a <= 'e' && a >= 'a') || (a <= 'E' && a >= 'A')
		if !b {
			return false
		}
	}
	return true
}
