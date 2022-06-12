/*
 * @lc app=leetcode id=1695 lang=golang
 *
 * [1695] Maximum Erasure Value
 */
package main

// @lc code=start
func maximumUniqueSubarray(nums []int) int {
	maxScore, score, begin, end := 0, 0, 0, 0
	m := make(map[int]bool)
	for ; end < len(nums); end++ {
		if m[nums[end]] {
			for ; true; begin++ {
				score -= nums[begin]
				delete(m, nums[begin])
				if nums[begin] == nums[end] {
					begin++
					break
				}
			}
		}
		score += nums[end]
		m[nums[end]] = true
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

// @lc code=end

// Accepted
// 62/62 cases passed (184 ms)
// Your runtime beats 92 % of golang submissions
// Your memory usage beats 96 % of golang submissions (8.8 MB)
