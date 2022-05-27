/*
 * @lc app=leetcode id=1342 lang=golang
 *
 * [1342] Number of Steps to Reduce a Number to Zero
 */
package leetcode

// @lc code=start
func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	numStep := 0
	for num != 0 {
		numStep += num&1 + 1
		num >>= 1
	}
	return numStep - 1
}

// @lc code=end

// Accepted
// 204/204 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 83.9 % of golang submissions (1.9 MB)
