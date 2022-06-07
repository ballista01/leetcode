/*
 * @lc app=leetcode id=448 lang=golang
 *
 * [448] Find All Numbers Disappeared in an Array
 */
package main

// @lc code=start
func findDisappearedNumbers(nums []int) []int {
	hash := make([]bool, len(nums)+1)
	for _, v := range nums {
		hash[v] = true
	}
	res := make([]int, 0)
	for i := 1; i < len(hash); i++ {
		if !hash[i] {
			res = append(res, i)
		}
	}
	return res
}

// @lc code=end

// Accepted
// 33/33 cases passed (63 ms)
// Your runtime beats 60.41 % of golang submissions
// Your memory usage beats 72.73 % of golang submissions (7.6 MB)
