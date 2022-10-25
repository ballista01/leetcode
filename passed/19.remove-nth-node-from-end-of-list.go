/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	index, last, pendingRemove := 0, -2, false
	removeHelper(head, n, index, &last, &pendingRemove)
	if pendingRemove {
		return head.Next
	}
	return head
}

func removeHelper(head *ListNode, n, index int, last *int, pendingRemove *bool) {
	if head == nil {
		*last = index - 1
		return
	}

	removeHelper(head.Next, n, index+1, last, pendingRemove)

	if *pendingRemove {
		head.Next = head.Next.Next
		*pendingRemove = false
	}
	if index == *last+1-n {
		*pendingRemove = true
	}
}

// @lc code=end

// Accepted
// 208/208 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 33.71 % of golang submissions (2.2 MB)
