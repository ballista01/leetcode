/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 */
package main

import "strings"

// @lc code=start
func solveNQueens(n int) [][]string {
	if n == 0 {
		return make([][]string, 0)
	}
	resStr := make([][]string, 0, n)
	board := make([][]bool, n)
	for i := range board {
		board[i] = make([]bool, n)
	}
	DFS(board, 0, &resStr)
	return resStr
}

func DFS(board [][]bool, rowIndex int, res *[][]string) {
	// all queens are placed, write solution to res
	if rowIndex == len(board) {
		strArr := make([]string, len(board))
		var sb strings.Builder
		for rowIndex, row := range board {
			sb.Grow(len(row))
			for _, v := range row {
				if v {
					sb.WriteRune('Q')
				} else {
					sb.WriteRune('.')
				}
			}
			strArr[rowIndex] = sb.String()
			sb.Reset()
		}
		*res = append(*res, strArr)
		return
	}

	// place the queen on board[rowIndex]
	for j := range board[rowIndex] {
		canAttack := false

		// check column
		for prevRowIndex := 0; prevRowIndex < rowIndex; prevRowIndex++ {
			if board[prevRowIndex][j] {
				canAttack = true
				break
			}
		}
		if canAttack {
			continue
		}

		// check descending diagonal
		for x, y := rowIndex, j; x >= 0 && y >= 0; x, y = x-1, y-1 {
			if board[x][y] {
				canAttack = true
				break
			}
		}
		if canAttack {
			continue
		}

		// check ascending diagonal
		for x, y := rowIndex, j; x >= 0 && y < len(board[x]); x, y = x-1, y+1 {
			if board[x][y] {
				canAttack = true
				break
			}
		}
		if canAttack {
			continue
		}

		board[rowIndex][j] = true
		DFS(board, rowIndex+1, res)
		board[rowIndex][j] = false
	}
}

// @lc code=end

// Accepted
// 9/9 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 82.77 % of golang submissions (3.1 MB)
