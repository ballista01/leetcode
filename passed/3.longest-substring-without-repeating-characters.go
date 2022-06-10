/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLen, nowLen := 0, 0
	rMap := make(map[rune]bool)
	for i, r := range s {
		if !rMap[r] {
			nowLen++
			rMap[r] = true
			if nowLen > maxLen {
				maxLen = nowLen
			}
		} else {
			nowLen++
			for j := i - nowLen + 1; j < i; j++ {
				rPre := rune(s[j])
				if rPre != r {
					delete(rMap, rPre)
					nowLen--
				} else {
					nowLen--
					break
				}
			}
		}
	}
	return maxLen
}

// @lc code=end

// Accepted
// 987/987 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 85.4 % of golang submissions (2.5 MB)
