package jianzhioffer

import "fmt"

// https://leetcode.cn/problems/qiu-12n-lcof/
func Run64() {
	fmt.Println(sumNums(10000))
}

func sumNums(n int) int {
	ans := 0
	A := n
	B := n + 1

	addGreatZero := func() bool {
		ans += A
		return ans > 0
	}

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	_ = ((B & 1) > 0) && addGreatZero()
	A <<= 1
	B >>= 1

	return ans >> 1
}
