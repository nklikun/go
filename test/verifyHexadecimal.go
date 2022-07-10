package test

import (
	"fmt"
)

func Run_verifyHexadecimal() {
	// 长度为1字节：10xx 10xx
	// 长度为2字节：110x 10xx 10xx 10xx
	// 长度为3字节：1110 10xx 10xx 10xx 10xx 10xx

	// s := BinToHexString("111010111001100010101001")
	// s := BinToHexString("10011011")
	s := test_BinToHexString("11100011")
	// s := "EA90F8"
	fmt.Println(verifyHex(s))
}

func test_BinToHexString(b string) string {
	s := ""
	for i := 0; i < len(b); i += 4 {
		n := 0
		for j := i; j < i+4; j++ {
			n = n << 1
			if b[j] == '1' {
				n++
			}
		}
		if n >= 0 && n <= 9 {
			s += string('0' + n)
		} else {
			s += string('A' + (n - 10))
		}
	}
	return s
}

// func test_intToBin(num int) string {
// 	s := ""
// 	for i := num; i > 0; i /= 2 {
// 		s = fmt.Sprintf("%d", i%2) + s
// 	}
// 	return s
// }

func verifyHex(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		oneCnt := len(s)/2 - i*4
		if oneCnt < 0 {
			oneCnt = 1
		}
		num := byteToi(s[i])
		for j := 3; j >= 0; j-- {
			if oneCnt > 0 {
				if 1&(num>>j) != 1 {
					return -1
				}
			} else if oneCnt == 0 {
				if 1&(num>>j) != 0 {
					return -1
				}
			} else {
				total = total << 1
				total += 1 & (num >> j)
			}
			oneCnt--
		}
	}
	return total
}

func byteToi(c byte) int {
	if c >= '0' && c <= '9' {
		return int(c - '0')
	}
	if c >= 'A' && c <= 'F' {
		return int(10 + c - 'A')
	}
	return -1
}
