/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */
package main

// @lc code=start
func search(nums []int, target int) int {
	breakpoint := findBreakpoint(nums)
	res1, res2 := binarySearch(nums, 0, breakpoint, target), binarySearch(nums, breakpoint, len(nums), target)
	nowPos := -1
	if res1 >= 0 {
		nowPos = res1
	} else {
		nowPos = res2
	}
	return nowPos
}
func findBreakpoint(nums []int) int {
	start, end := 0, len(nums)-1
	var mid, br int
	for start < end {
		mid = (start + end) / 2
		if nums[mid] > nums[start] {
			if nums[mid+1] < nums[start] {
				br = mid + 1
				break
			}
			start = mid
		} else if nums[mid] < nums[start] {
			end = mid
			if nums[mid-1] > nums[mid] {
				br = mid
				break
			}
			end = mid
		} else {
			if nums[mid] < nums[end] {
				br = 0
				break
			} else {
				br = end
				break
			}

		}
	}
	return br
}
func binarySearch(nums []int, start int, end int, target int) int {
	end--
	var mid int
	for start <= end {
		mid = (start + end) / 2
		if target < nums[mid] {
			end = mid - 1
		} else if target > nums[mid] {
			start = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end

// Accepted
// 195/195 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 63.09 % of golang submissions (2.6 MB)
