/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */
package leetcode

// @lc code=start
func missingNumber(nums []int) int {
	totalLen := len(nums) + 1
	totalSum := (totalLen - 1) * totalLen / 2
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return totalSum - sum
}

// @lc code=end

// Accepted
// 122/122 cases passed (16 ms)
// Your runtime beats 82.97 % of golang submissions
// Your memory usage beats 97.21 % of golang submissions (6.2 MB)
