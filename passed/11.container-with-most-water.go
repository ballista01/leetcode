/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */
package main

// @lc code=start
func maxArea(height []int) int {
	vol, maxVol := 0, 0
	front, back := 0, len(height)-1
	for front < back {
		if height[back] < height[front] {
			vol = height[back] * (back - front)
			back--
		} else {
			vol = height[front] * (back - front)
			front++
		}
		if vol > maxVol {
			maxVol = vol
		}
	}
	return maxVol
}

// @lc code=end

// Accepted
// 60/60 cases passed (69 ms)
// Your runtime beats 99.88 % of golang submissions
// Your memory usage beats 82.09 % of golang submissions (8.7 MB)
