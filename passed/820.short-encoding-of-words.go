/*
 * @lc app=leetcode id=820 lang=golang
 *
 * [820] Short Encoding of Words
 */
package main

// @lc code=start
func minimumLengthEncoding(words []string) int {
	m := make(map[string]byte, len(words)*7)
	tLen := 0
	for _, word := range words {
		wKey, exists := m[word]
		if !exists || wKey == 0 {
			m[word] = 2
			wLen := len(word)
			tLen += (wLen + 1)
			for i := 1; i < wLen; i++ {
				sub := word[i:wLen]
				sKey, exists := m[sub]
				if exists && sKey == 2 {
					tLen -= (len(sub) + 1)
				}
				m[sub] = 1
			}
		}
	}
	return tLen
}

// @lc code=end

// Accepted
// 36/36 cases passed (36 ms)
// Your runtime beats 47.06 % of golang submissions
// Your memory usage beats 52.94 % of golang submissions (7.5 MB)
