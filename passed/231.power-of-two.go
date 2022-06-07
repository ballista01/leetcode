/*
 * @lc app=leetcode id=231 lang=golang
 *
 * [231] Power of Two
 */
package main

// @lc code=start
func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for ; n > 0; n >>= 1 {
		if n&1 == 1 && n>>1 > 0 {
			return false
		}
	}
	return true
}

// @lc code=end

// Accepted
// 1108/1108 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 73.03 % of golang submissions (2 MB)
