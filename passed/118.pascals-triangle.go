/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */
package main

// @lc code=start
func generate(numRows int) [][]int {
	res := make([][]int, 0, numRows)
	res = append(res, []int{1})
	for i := 1; i < numRows; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[len(row)-1] = 1
		for j := 1; j < len(row)-1; j++ {
			row[j] = res[i-1][j-1] + res[i-1][j]
		}
		res = append(res, row)
	}
	return res
}

// @lc code=end

// Accepted
// 14/14 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 55.81 % of golang submissions (2.1 MB)
