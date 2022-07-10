package main

import "fmt"

// https://leetcode.cn/problems/sudoku-solver/
func run37() {
	board := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	solveSudoku(board)
	run37_print(board)
}

func run37_print(board [][]byte) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("'")
			fmt.Printf(string(board[i][j]))
			fmt.Printf("',")
		}
		fmt.Println()
	}
	fmt.Println()
}

func solveSudoku(board [][]byte) {
	hor := make([]int, 9)
	ver := make([]int, 9)
	squ := make([]int, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - 48)
				hor[i] += 1 << num
				ver[j] += 1 << num
				k := (i/3)*3 + j/3
				squ[k] += 1 << num
			}
		}
	}

	board, _ = solveSudoku_helper(hor, ver, squ, board, 0, 0)
}

// return bool: found the solution, no need more explorer
func solveSudoku_helper(hor, ver, squ []int, board [][]byte, i, j int) ([][]byte, bool) {
	i, j = solveSudoku_getNextPostion(board, i, j)
	// full filled
	if i == -1 {
		return board, true
	}

	for n := 1; n < 10; n++ {
		if hor[i]&(1<<n) == 0 &&
			ver[j]&(1<<n) == 0 &&
			squ[(i/3)*3+j/3]&(1<<n) == 0 {
			board[i][j] = byte(n + 48)
			ver[j] += 1 << n
			hor[i] += 1 << n
			squ[(i/3)*3+j/3] += 1 << n
			if _, b := solveSudoku_helper(hor, ver, squ, board, i, j); b {
				return nil, true
			}

			// remove the filled, try other numbers
			// recover board, becasue solveSudoku_getNextPostion needs,
			// otherwise, we can cover it with new numbers since ver/hor/squ setted
			board[i][j] = '.'
			ver[j] -= 1 << n
			hor[i] -= 1 << n
			squ[(i/3)*3+j/3] -= 1 << n
		}
	}
	return nil, false
}

func solveSudoku_getNextPostion(board [][]byte, i, j int) (int, int) {
	for i < 9 {
		if j == 9 {
			i++
			j = 0
			continue
		}
		if board[i][j] == '.' {
			return i, j
		}
		j++
	}
	return -1, j
}

/*
func solveSudoku(board [][]byte) {
	hor := make([]int, 9)
	ver := make([]int, 9)
	squ := make([]int, 9)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				num := int(board[i][j] - 48)
				hor[i] += 1 << num
				ver[j] += 1 << num
				k := (i/3)*3 + j/3
				squ[k] += 1 << num
			}
		}
	}

	// return: found the solution, no need more explore
	var helper func([]int, []int, []int, [][]byte, int, int) bool
	helper = func(hor, ver, squ []int, b [][]byte, i, j int) bool {
		i, j = solveSudoku_getNextPostion(b, i, j)
		// full filled
		if i == -1 {
			// board = b
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					board[i][j] = b[i][j]
				}
			}
			return true
		}
		// run37_print(b)

		for n := 1; n < 10; n++ {
			if hor[i]&(1<<n) == 0 &&
				ver[j]&(1<<n) == 0 &&
				squ[(i/3)*3+j/3]&(1<<n) == 0 {
				b[i][j] = byte(n + 48)
				ver[j] += 1 << n
				hor[i] += 1 << n
				squ[(i/3)*3+j/3] += 1 << n
				if helper(hor, ver, squ, b, i, j) {
					return true
				}

				// remove the filled, try other numbers
				// recover board, becasue solveSudoku_getNextPostion needs,
				// otherwise, we can cover it with new numbers since ver/hor/squ setted
				b[i][j] = '.'
				ver[j] -= 1 << n
				hor[i] -= 1 << n
				squ[(i/3)*3+j/3] -= 1 << n
			}
		}
		return false
	}
	helper(hor, ver, squ, board, 0, 0)
}

func solveSudoku_getNextPostion(board [][]byte, i, j int) (int, int) {
	for i < 9 {
		if j == 9 {
			i++
			j = 0
			continue
		}
		if board[i][j] == '.' {
			return i, j
		}
		j++
	}
	return -1, j
}
*/
