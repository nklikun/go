package main

import (
	"fmt"
)

/*
592. 分数加减运算
给定一个表示分数加减运算的字符串 expression ，你需要返回一个字符串形式的计算结果。
这个结果应该是不可约分的分数，即最简分数。 如果最终结果是一个整数，例如 2，你需要将它转换成分数形式，其分母为 1。所以在上述例子中, 2 应该被转换为 2/1。

示例 1:
输入: expression = "-1/2+1/2"
输出: "0/1"

示例 2:
输入: expression = "-1/2+1/2+1/3"
输出: "1/3"

示例 3:
输入: expression = "1/3-1/2"
输出: "-1/6"

提示:
输入和输出字符串只包含 '0' 到 '9' 的数字，以及 '/', '+' 和 '-'。
输入和输出分数格式均为 ±分子/分母。如果输入的第一个分数或者输出的分数是正数，则 '+' 会被省略掉。
输入只包含合法的最简分数，每个分数的分子与分母的范围是  [1,10]。 如果分母是1，意味着这个分数实际上是一个整数。
输入的分数个数范围是 [1,10]。
最终结果的分子与分母保证是 32 位整数范围内的有效整数。

https://leetcode.cn/problems/fraction-addition-and-subtraction/
*/
func Run592() {
	fmt.Println(fractionAddition("-1/2+1/2+1/3+1/2+1+12/13+5/26"))
}

func fractionAddition(expression string) string {
	numerator, denominator := 0, 1
	// 临时存储分子分母和符号
	nu, de, neg := 0, 0, false
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	add := func() {
		if nu == 0 && de == 0 {
			return
		}
		if neg {
			nu = -nu
		}
		if de == 0 {
			de = 1
		}
		//分母相乘，并计算分子
		denominator, numerator = denominator*de, numerator*de+denominator*nu
		if numerator != 0 {
			//进行约分
			n1, n2 := abs(numerator), abs(denominator)
			if n1 < n2 {
				n1, n2 = n2, n1
			}
			for n1%n2 != 0 {
				n1, n2 = n2, n1%n2
			}
			denominator /= n2
			numerator /= n2
		} else {
			numerator, denominator = 0, 1
		}
		nu, de, neg = 0, 0, false
	}
	//逐字符读取，如遇到+-则进行计算
	for i := 0; i < len(expression); i++ {
		if expression[i] == '+' || expression[i] == '-' {
			add()
			neg = expression[i] == '-'
		} else {
			//一次性读取完整一个分数，如-12/35
			//先读取分子，遇到/则读取分母，直到结束或者+-
			slash := false
			for ; i < len(expression); i++ {
				if expression[i] == '+' || expression[i] == '-' {
					i--
					break
				}
				if expression[i] == '/' {
					slash = true
					continue
				}
				if !slash {
					nu *= 10
					nu += int(expression[i] - '0')
				} else {
					de *= 10
					de += int(expression[i] - '0')
				}
			}
		}
	}
	add()
	return fmt.Sprintf("%d/%d", numerator, denominator)
}
