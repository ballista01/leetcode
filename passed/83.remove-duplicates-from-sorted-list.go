/*
 * @lc app=leetcode id=83 lang=golang
 *
 * [83] Remove Duplicates from Sorted List
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	var prev, now *ListNode = head, head.Next
	for now != nil {
		if now.Val == prev.Val {
			for now != nil && now.Val == prev.Val {
				now = now.Next
			}
			prev.Next = now
		} else {
			prev = prev.Next
			if prev == nil {
				break
			} else {
				now = prev.Next
			}
		}

	}
	return head
}

// @lc code=end

// Accepted
// 166/166 cases passed (3 ms)
// Your runtime beats 83.92 % of golang submissions
// Your memory usage beats 86.89 % of golang submissions (3 MB)
