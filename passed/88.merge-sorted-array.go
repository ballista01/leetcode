/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
package main

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	m--
	n--
	index := len(nums1) - 1
	for {
		if m >= 0 && n >= 0 {
			if nums1[m] > nums2[n] {
				nums1[index] = nums1[m]
				m--
			} else {
				nums1[index] = nums2[n]
				n--
			}
			index--
		} else if m >= 0 && n < 0 {
			nums1[index] = nums1[m]
			m--
			index--
		} else if m < 0 && n >= 0 {
			nums1[index] = nums2[n]
			n--
			index--
		} else {
			break
		}
	}
}

// @lc code=end

// Accepted
// 59/59 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 83.25 % of golang submissions (2.3 MB)
