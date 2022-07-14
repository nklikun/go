package main

//https://leetcode.cn/problems/integer-replacement/
func run397() {

}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}

	if n%2 == 0 {
		return integerReplacement(n/2) + 1
	} else {
		a := integerReplacement((n - 1) / 2)
		b := integerReplacement((n + 1) / 2)
		if a > b {
			return b + 2
		} else {
			return a + 2
		}
	}
}
