/*
 * @lc app=leetcode id=318 lang=golang
 *
 * [318] Maximum Product of Word Lengths
 */
package leetcode

import "sort"

// @lc code=start
func maxProduct(words []string) int {
	maxVal := 0
	wordBits := make([]uint, len(words))
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})
	for i, w := range words {
		for _, r := range w {
			wordBits[i] |= (0x1 << uint(r-'a'))
		}
	}
	for i := 0; i < len(words)-1; i++ {
		if maxVal > len(words[i])*len(words[i+1]) {
			break
		}
		for j := i + 1; j < len(words); j++ {
			nowVal := len(words[i]) * len(words[j])
			if nowVal <= maxVal || (wordBits[i]&wordBits[j] > 0) {
				continue
			} else {
				maxVal = nowVal
			}
		}
	}
	return maxVal
}

// @lc code=end

// Accepted
// 167/167 cases passed (11 ms)
// Your runtime beats 83.33 % of golang submissions
// Your memory usage beats 93.75 % of golang submissions (6.3 MB)
