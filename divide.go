package main

import (
	"fmt"
	"math"
)

// https://leetcode.cn/problems/divide-two-integers
func run29() {
	fmt.Println(divide(-2147483648, -1))
}

func divide(dividend int, divisor int) int {
	MAX_INT := math.MaxInt32
	MIN_INT := math.MinInt32
	if dividend == 0 {
		return 0
	}
	if dividend == MIN_INT {
		if divisor == 1 {
			return MIN_INT
		}
		if divisor == -1 {
			return MAX_INT
		}
	}
	if divisor == MIN_INT {
		if dividend == MIN_INT {
			return 1
		}
		return 0
	}

	negative := (dividend ^ divisor) < 0
	t := int64(dividend)
	if t < 0 {
		t = ^t + 1
	}
	d := int64(divisor)
	if d < 0 {
		d = ^d + 1
	}

	n := 0
	for t >= d {
		i := 0
		for ; (d << i) <= t; i++ {
		}
		n += (1 << (i - 1))
		t = t - (d << (i - 1))
	}

	if negative {
		return ^n + 1
	}
	return n
}
