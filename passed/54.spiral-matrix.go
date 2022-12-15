/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	dir := [][]int{}
	dir = append(dir, []int{0, 1})
	dir = append(dir, []int{1, 0})
	dir = append(dir, []int{0, -1})
	dir = append(dir, []int{-1, 0})

	maxRow, maxCol := len(matrix), len(matrix[0])
	var nextRow, nextCol int
	res := make([]int, maxRow*maxCol)
	numFilled := 0
	row, col := 0, 0
	dirIndex := 0
	for numFilled < maxRow*maxCol {
		res[numFilled] = matrix[row][col]
		numFilled++
		// fmt.Printf("row=%d, col=%d, val=%d, ", row, col, matrix[row][col])
		matrix[row][col] = 400
		nextRow, nextCol = row+dir[dirIndex][0], col+dir[dirIndex][1]
		if nextCol >= maxCol || nextRow >= maxRow || nextCol < 0 || nextRow < 0 || matrix[nextRow][nextCol] > 100 {
			dirIndex++
			dirIndex %= len(dir)
			nextRow, nextCol = row+dir[dirIndex][0], col+dir[dirIndex][1]
			row, col = nextRow, nextCol
			// fmt.Printf("REDIR: nextRow=%d, nextCol=%d\n", nextRow, nextCol)
			continue
		}
		row, col = nextRow, nextCol
		// fmt.Printf("nextRow=%d, nextCol=%d\n", nextRow, nextCol)
	}
	return res
}

// @lc code=end

// Accepted
// 23/23 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 74.64 % of golang submissions (2 MB)
