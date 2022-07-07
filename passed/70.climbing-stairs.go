/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */
package main

// @lc code=start
var hash [46]int

func climbStairs(n int) int {
	if n < 0 {
		return 0
	}
	if n == 0 {
		return 1
	}
	if hash[n] > 0 {
		return hash[n]
	} else {
		hash[n] = climbStairs(n-1) + climbStairs(n-2)
		return hash[n]
	}
}

// @lc code=end

// Accepted
// 45/45 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 90.44 % of golang submissions (1.9 MB)
