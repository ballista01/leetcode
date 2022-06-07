/*
 * @lc app=leetcode id=88 lang=golang
 *
 * [88] Merge Sorted Array
 */
package pending

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	x, y, xe := m-1, n-1, m+n-1
	for x >= 0 && y >= 0 {
		if nums1[x] >= nums2[y] {
			nums1[xe] = nums1[x]
			xe--
			x--
		} else {
			nums1[xe] = nums2[y]
			xe--
			y--
		}
	}
	for ; y >= 0; y-- {
		nums1[xe] = nums2[y]
		xe--
	}

}

// @lc code=end

// Accepted
// 59/59 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 76.03 % of golang submissions (2.2 MB)
