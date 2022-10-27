/*
 * @lc app=leetcode id=278 lang=golang
 *
 * [278] First Bad Version
 */
package main

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	lo, mid, hi := 1, 0, n
	for lo <= hi {
		mid = (lo + hi) / 2
		if isBadVersion(mid) {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end

// Accepted
// 24/24 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 87.17 % of golang submissions (1.9 MB)
