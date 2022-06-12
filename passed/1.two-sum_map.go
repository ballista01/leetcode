/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, v := range nums {
		comp, existComp := m[target-v]
		if existComp {
			return []int{comp, i}
		} else {
			m[v] = i
		}
	}
	return nil
}

// @lc code=end

// Accepted
// 57/57 cases passed (4 ms)
// Your runtime beats 93.49 % of golang submissions
// Your memory usage beats 52.23 % of golang submissions (4.2 MB)
