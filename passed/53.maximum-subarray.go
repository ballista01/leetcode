/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	now, max := 0, nums[0]
	for _, v := range nums {
		now += v
		if now > max {
			max = now
		}
		if now < 0 {
			now = 0
		}
	}
	return max
}

// @lc code=end

// Accepted
// 209/209 cases passed (141 ms)
// Your runtime beats 52.48 % of golang submissions
// Your memory usage beats 75.8 % of golang submissions (9.1 MB)
