package main

// https://leetcode.cn/problems/pascals-triangle
func run118() {
	generate(5)
}

func generate(numRows int) [][]int {
	triangle := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		nums := make([]int, i+1)
		nums[0] = 1
		nums[i] = 1
		if i > 0 {
			last := triangle[i-1]
			for j := 1; j < i; j++ {
				nums[j] = last[j-1] + last[j]
			}
		}
		triangle[i] = nums
	}

	return triangle
}
