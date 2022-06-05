/*
 * @lc app=leetcode id=867 lang=golang
 *
 * [867] Transpose Matrix
 */
package main

// @lc code=start
func transpose(matrix [][]int) [][]int {
	mt := make([][]int, len(matrix[0]))
	for j := range mt {
		mt[j] = make([]int, len(matrix))
		for i := 0; i < len(mt[j]); i++ {
			mt[j][i] = matrix[i][j]
		}
	}
	return mt
}

// @lc code=end

// Accepted
// 36/36 cases passed (10 ms)
// Your runtime beats 50.98 % of golang submissions
// Your memory usage beats 62.75 % of golang submissions (6.3 MB)
