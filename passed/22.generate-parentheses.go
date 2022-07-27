/*
 * @lc app=leetcode id=22 lang=golang
 *
 * [22] Generate Parentheses
 */
package main

import "bytes"

// @lc code=start
func generateParenthesis(n int) []string {
	resArr := make([]string, 0, n*n)
	comb := make([]byte, 0, n*2)

	generateParenthesisHelper(n*2, 0, comb, &resArr)
	return resArr
}

func generateParenthesisHelper(n int, balance int, comb []byte, resArr *[]string) {
	if n == 0 && balance == 0 {
		combStr := bytes.NewBuffer(comb).String()
		*resArr = append(*resArr, combStr)
		return
	}
	if balance > n || balance < 0 {
		return
	}

	comb = append(comb, '(')
	generateParenthesisHelper(n-1, balance+1, comb, resArr)
	comb[len(comb)-1] = ')'
	generateParenthesisHelper(n-1, balance-1, comb, resArr)
}

// @lc code=end

// Accepted
// 8/8 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 73.87 % of golang submissions (2.8 MB)
