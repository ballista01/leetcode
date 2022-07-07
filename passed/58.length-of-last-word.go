/*
 * @lc app=leetcode id=58 lang=golang
 *
 * [58] Length of Last Word
 */
package main

// @lc code=start
func lengthOfLastWord(s string) int {
	lastLen := 0
	for i := len(s) - 1; i >= 0; i-- {
		if lastLen == 0 {
			if s[i] != ' ' {
				lastLen++
			}
		} else {
			if s[i] == ' ' {
				break
			}
			lastLen++
		}
	}
	return lastLen
}

// @lc code=end

// Accepted
// 58/58 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 86.87 % of golang submissions (2 MB)
