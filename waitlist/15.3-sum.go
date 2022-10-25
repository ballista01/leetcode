/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	resMap := make(map[int]bool)
	res := make([][]int, 0, 10)
	sort.Ints(nums)
	if nums[0] >= 0 && nums[len(nums)-1] > 0 {
		return res
	}
	if nums[0] < 0 && nums[len(nums)-1] <= 0 {
		return res
	}
	for front, back := 0, len(nums)-1; front+2 <= back; {
		for mid := front + 1; mid < back; mid++ {
			if nums[front]+nums[mid]+nums[back] == 0 {
				hash := (0-nums[front])*10000000000 + nums[mid]*100000 + nums[back]
				_, combExist := resMap[hash]
				if !combExist {
					res = append(res, []int{nums[front], nums[mid], nums[back]})
				}
				resMap[hash] = true
			}
		}
		if 0-nums[front] > nums[back] {
			front++
		} else {
			back--
		}
	}
	return res
}

// @lc code=end
