/*
 * @lc app=leetcode id=884 lang=golang
 *
 * [884] Uncommon Words from Two Sentences
 */

package main

import "strings"

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) []string {
	m := make(map[string]int)
	w1 := strings.Split(s1, " ")
	w2 := strings.Split(s2, " ")
	for _, v := range w1 {
		_, exist := m[v]
		if exist {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range w2 {
		_, exist := m[v]
		if exist {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	res := make([]string, 0, 10)
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}

// @lc code=end

// Accepted
// 55/55 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 28.95 % of golang submissions (2.3 MB)
