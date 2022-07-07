/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */
package main

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	maxLen, nowLen := 0, 0
	hash := make(map[rune]int, 26)
	for i, v := range s {
		nowLen++
		lastIndex, collision := hash[v]
		if collision {
			for j := i - (nowLen - 1); j < lastIndex; j++ {
				delete(hash, rune(s[j]))
			}
			nowLen = i - lastIndex
			hash[v] = i
		} else {
			hash[v] = i
		}

		if nowLen > maxLen {
			maxLen = nowLen
		}
	}
	return maxLen
}

// @lc code=end

// Accepted
// 987/987 cases passed (3 ms)
// Your runtime beats 94.17 % of golang submissions
// Your memory usage beats 70.33 % of golang submissions (2.8 MB)
