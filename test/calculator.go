package test

import (
	"fmt"
	"strconv"
)

func RunCalculator() {
	// var s string
	// fmt.Scan(&s)
	// s := "9+(3-1)*((5+10)/3)" // 19
	// s := "3*5+8-0*3-6+0+0" // 17
	s := "9-1+0-9+4-10-(8+4+10)-1" // -30
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	if len(s) == 0 {
		return 0
	}
	nums := make([]int, 0)
	ops := make([]byte, 0)

	doCalc := func() {
		// 使用匿名函数避免切片传入拷贝到新地址,
		// 传入前的原指针不会指向新地址
		b := nums[len(nums)-1]
		a := nums[len(nums)-2]
		nums = nums[:len(nums)-1]

		op := ops[len(ops)-1]
		ops = ops[:len(ops)-1]

		if op == '+' {
			a = a + b
		} else if op == '-' {
			a = a - b
		} else if op == '*' {
			a = a * b
		} else if op == '/' {
			a = a / b
		}
		nums[len(nums)-1] = a
	}

	//保证计算到最开头
	if s[0] != '(' {
		ops = append(ops, '(')
		s += ")"
	}

	nextIsOp := false
	for i := 0; i < len(s); i++ {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			ops = append(ops, '(')
		} else if s[i] == ')' || s[i] == ']' || s[i] == '}' {
			for ops[len(ops)-1] != '(' {
				doCalc()
			}
			ops = ops[:len(ops)-1]
		} else if nextIsOp {
			for compareOps(ops[len(ops)-1], s[i]) {
				doCalc()
			}
			ops = append(ops, s[i])
			nextIsOp = false
		} else {
			j := i
			// 负号或者正号
			if s[j] == '-' || s[j] == '+' {
				i++
			}
			for s[i] <= '9' && s[i] >= '0' {
				i++
			}
			n, _ := strconv.Atoi(s[j:i])
			nums = append(nums, n)
			i--
			nextIsOp = true
		}
	}
	return nums[len(nums)-1]
}

func compareOps(a, b byte) bool {
	if a == '(' {
		return false
	}
	if (a == '+' || a == '-') && (b == '*' || b == '/') {
		return false
	}
	return true
}
