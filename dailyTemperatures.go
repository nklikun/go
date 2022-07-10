package main

import "fmt"

// https://leetcode.cn/problems/daily-temperatures/
func run739() {
	expect := []int{3, 1, 1, 4, 3, 1, 1, 3, 1, 1, 1, 1, 30, 1, 3, 2, 1, 25, 2, 1, 19, 2, 1, 3, 2, 1, 11, 5, 1, 1, 2, 1, 3, 2, 1, 2, 1, 2, 1, 3, 2, 1, 0, 46, 3, 1, 1, 1, 30, 18, 5, 1, 1, 2, 1, 12, 1, 10, 5, 1, 2, 1, 1, 4, 3, 1, 1, 11, 1, 1, 8, 1, 1, 5, 1, 3, 1, 1, 11, 1, 3, 2, 1, 1, 5, 3, 2, 1, 1, 0, 1, 0, 3, 2, 1, 0, 0, 2, 1, 0}
	temps := []int{64, 40, 49, 73, 72, 35, 68, 83, 35, 73, 84, 88, 96, 43, 74, 63, 41, 95, 48, 46, 89, 72, 34, 85, 72, 59, 87, 49, 30, 32, 47, 34, 74, 58, 31, 75, 73, 88, 64, 92, 83, 64, 100, 99, 81, 41, 48, 83, 96, 92, 82, 32, 35, 68, 68, 92, 73, 92, 52, 33, 44, 38, 47, 88, 71, 50, 57, 95, 33, 65, 94, 44, 47, 79, 41, 74, 50, 67, 97, 31, 68, 50, 37, 70, 77, 55, 48, 30, 77, 100, 31, 100, 69, 60, 47, 95, 68, 47, 33, 64}
	// expect := []int{1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0}
	// temps := []int{30, 31, 32, 33, 34, 35, 30, 31, 32, 33, 34, 35}
	fmt.Println(expect)
	result := dailyTemperatures(temps)
	// fmt.Println(result)
	if len(result) != len(expect) {
		fmt.Println("false, length error")
		return
	}
	fmt.Printf("[")
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
		if result[i] != expect[i] {
			fmt.Printf("\nfalse: index: %d, temperatures: %d", i, temps[i])
			return
		}
	}
	fmt.Printf("]\n")
	fmt.Println(true)
}

func dailyTemperatures(temperatures []int) []int {
	if len(temperatures) == 0 {
		return nil
	}
	stack, ans := make([]int, 0), make([]int, len(temperatures))
	for i := 0; i < len(temperatures); i++ {
		if len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			for len(stack) > 0 {
				// k : stored last index of temperatures
				k := stack[len(stack)-1]
				if temperatures[k] < temperatures[i] {
					ans[k] = i - stack[len(stack)-1]
					stack = stack[:len(stack)-1]
				} else {
					break
				}
			}
		}
		stack = append(stack, i)
	}
	return ans
}
