/*
 * @lc app=leetcode id=67 lang=golang
 *
 * [67] Add Binary
 */
package main

import "bytes"

// @lc code=start
func addBinary(a string, b string) string {
	resLen := len(a)
	if len(b) > len(a) {
		resLen = len(b)
	}
	resLen++
	arr := make([]byte, resLen)

	k := resLen - 1
	var carry byte = 0
	var x, y, digi byte
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0; i, j, k = i-1, j-1, k-1 {
		if i >= 0 {
			x = a[i] - '0'
		} else {
			x = 0
		}
		if j >= 0 {
			y = b[j] - '0'
		} else {
			y = 0
		}

		digi = (carry + x + y) % 2
		carry = (carry + x + y) / 2
		arr[k] = digi + '0'
	}
	if carry == 0 {
		return bytes.NewBuffer(arr[1:]).String()
	} else {
		arr[k] = carry + '0'
		return bytes.NewBuffer(arr).String()
	}
}

// @lc code=end

// Accepted
// 294/294 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 94.23 % of golang submissions (2.2 MB)
