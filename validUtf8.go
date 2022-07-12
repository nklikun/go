package main

import "fmt"

// https://leetcode.cn/problems/utf-8-validation/
func run393() {
	fmt.Println(validUtf8([]int{197, 130, 1, 235, 140, 130}))
	fmt.Println(validUtf8([]int{197, 130, 1, 235, 140, 4})) // 4 is not vaild
	fmt.Println(validUtf8([]int{248, 130, 130, 130}))       // length > 4 && data length < length
	fmt.Println(validUtf8([]int{250, 145, 145, 145, 145}))  // length > 4
	fmt.Println(validUtf8([]int{0b10011001}))               // length is 1， not 2-4
}

func test_printBin(num ...int) {
	for _, n := range num {
		s := ""
		for i := n; i > 0; i /= 2 {
			s = fmt.Sprintf("%d", i%2) + s
		}
		for i := len(s); i < 8; i++ {
			s = "0" + s
		}
		fmt.Printf(s + ",")
	}
}

func validUtf8(data []int) bool {
	m := 0b10
	l := 0
	for l < len(data) {
		// for 1 byte
		if (data[l]>>7)&1 == 0 {
			l = l + 1
			continue
		}

		// for 2-4 bytes
		// get length of next several bytes
		r := l
		for i := 7; i >= 0; i-- {
			if (data[l]>>i)&1 == 0 {
				break
			}
			r++
		}
		// the left is not enough as the first byte tells
		if r-l < 2 || r-l > 4 || len(data) < r {
			return false
		}
		for i := l + 1; i < r; i++ {
			if data[i]>>6 != m {
				return false
			}
		}
		l = r
	}
	return true
}
