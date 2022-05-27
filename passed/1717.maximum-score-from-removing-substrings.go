/*
 * @lc app=leetcode id=1717 lang=golang
 *
 * [1717] Maximum Score From Removing Substrings
 */
package leetcode

import (
	"container/list"
)

// @lc code=start
func maximumGain(s string, x int, y int) int {
	var (
		big, small       = x, y
		bigStr, smallStr = "ab", "ba"
		score            = 0
	)
	if y > x {
		big, bigStr, small, smallStr = small, smallStr, big, bigStr
	}
	bigStack, smallStack := list.New(), list.New()

	for i, _ := range s {
		if bigStack.Len() != 0 && bigStack.Back().Value == bigStr[0] && s[i] == bigStr[1] {
			score += big
			bigStack.Remove(bigStack.Back())
		} else {
			bigStack.PushBack(s[i])
		}
	}
	for bigStack.Len() != 0 {
		if smallStack.Len() != 0 && bigStack.Back().Value == smallStr[0] && smallStack.Back().Value == smallStr[1] {
			score += small
			smallStack.Remove(smallStack.Back())
		} else {
			smallStack.PushBack(bigStack.Back().Value)
		}
		bigStack.Remove(bigStack.Back())
	}
	return score
}

// @lc code=end

// Accepted
// 76/76 cases passed (342 ms)
// Your runtime beats 66.67 % of golang submissions
// Your memory usage beats 33.33 % of golang submissions (11.6 MB)
