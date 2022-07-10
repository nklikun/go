package main

import "fmt"

// https://leetcode.cn/problems/counting-bits
func run338() {
	// 	18/2 = 9... 0 .. 18
	// 	9/2 = 4 ... 1 .. 9
	// 	4/2 = 2 ... 0 .. 4
	// 	2/2 = 1 ... 0 .. 2
	//  1/2 = 0 ... 1

	// 	18 = 10010  = 2
	// 	17 = 10001  = 2
	// 	16 = 10000  = 1
	// 	15 = 01111  = 4
	fmt.Println(countBits(18))
}

/*func countBits(n int) []int {
	counts := make([]int, n+1)
	high := 0
	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			high = i
		}
		counts[i] = counts[i-high] + 1
	}
	return counts
}*/

func countBits(n int) []int {
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		counts[i] = counts[i>>1] + i&1
	}
	return counts
}

// func count(n int) {

// 	14 =

// }
