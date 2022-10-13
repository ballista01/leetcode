/*
 * @lc app=leetcode id=119 lang=golang
 *
 * [119] Pascal's Triangle II
 */
package main

// @lc code=start
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	res[0], res[1] = 1, 1
	if rowIndex == 1 {
		return res
	}

	upper := 1
	for i := 1; i < rowIndex; i++ {
		upper++
		res[upper] = 1
		for j := upper - 1; j > 0; j-- {
			res[j] = res[j-1] + res[j]
		}
	}
	return res
}

// @lc code=end

// Accepted
// 34/34 cases passed (1 ms)
// Your runtime beats 58.84 % of golang submissions
// Your memory usage beats 100 % of golang submissions (1.9 MB)
