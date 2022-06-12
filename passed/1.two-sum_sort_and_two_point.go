/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */
package main

import "sort"

func main() { twoSum([]int{3, 2, 4}, 6) }

// @lc code=start
func twoSum(nums []int, target int) []int {
	arr := make([]IndexedNum, len(nums))
	for i, v := range nums {
		arr[i].val = v
		arr[i].index = i
	}
	sort.Slice(arr, func(p, q int) bool { return arr[p].val < arr[q].val })
	sum := 0
	for begin, end := 0, len(arr)-1; begin < end; {
		sum = arr[begin].val + arr[end].val
		if sum < target {
			begin++
		} else if sum > target {
			end--
		} else {
			return []int{arr[begin].index, arr[end].index}
		}
	}
	return nil
}

type IndexedNum struct {
	index, val int
}

// @lc code=end

// Accepted
// 57/57 cases passed (8 ms)
// Your runtime beats 73.74 % of golang submissions
// Your memory usage beats 65.86 % of golang submissions (4.1 MB)
