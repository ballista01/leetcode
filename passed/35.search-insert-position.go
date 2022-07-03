/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */
package main

// @lc code=start
func searchInsert(nums []int, target int) int {
	begin, end := 0, len(nums)-1
	for begin < end {
		mid := (begin + end) / 2
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			begin = mid + 1
		} else {
			return mid
		}
	}
	if nums[begin] < target {
		return begin + 1
	} else {
		return begin
	}
}

// @lc code=end

// Accepted
// 64/64 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 75.86 % of golang submissions (3 MB)
