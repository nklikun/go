package main

import "fmt"

// https://leetcode.cn/problems/valid-sudoku
func run36() {
	// a := uint(0)
	// a += 1 << 2
	// a += 1 << 3
	// fmt.Println(a)
	// fmt.Println(a & (1 << 3))
	// fmt.Println(a & (1 << 2))

	// num := int('1' - 48)
	// fmt.Println(num)
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	// board := [][]byte{
	// 	{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
	// 	{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
	// 	{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
	// 	{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
	// 	{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
	// 	{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
	// 	{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}

	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	// 划分为三种map类型map[k]v，k规则如下。v规则:以位存储，如数字8， v=1<<8，对比也按位对比，
	// '.'记作0。
	// 第一种：九宫格squared，其角标来自于uint(x/3 + 1),uint(y/3 + 1),
	//        squ[k]v, k = uint(x/3 + 1)*10 + uint(y/3 + 1)
	// 第二种：横坐标horizontal，角标来自于y, hor[k]v, k = y
	// 第三种：纵坐标vertical，角标来自于x, ver[k]v, k = x

	squ := make(map[uint]uint, 9)
	hor := make(map[uint]uint, 9)
	ver := make(map[uint]uint, 9)
	for i := uint(0); i < 9; i++ {
		for j := uint(0); j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := int(board[i][j] - 48)

			if v, ok := hor[i]; ok {
				if v&(1<<num) != 0 {
					return false
				}
			}
			hor[i] += (1 << num)

			if v, ok := ver[j]; ok {
				if v&(1<<num) != 0 {
					return false
				}
			}
			ver[j] += (1 << num)

			k := uint(i/3+1)*10 + uint(j/3+1)
			if v, ok := squ[k]; ok {
				if v&(1<<num) != 0 {
					return false
				}
			}
			squ[k] += (1 << num)
			// fmt.Printf("i:%d j:%d\n", i, j)
		}
	}
	return true
}
