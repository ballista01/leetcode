/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	root := new(ListNode)
	prev := root
	var now *ListNode
	var v1, v2 int
	for l1 != nil || l2 != nil {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
			l2 = l2.Next
		}

		now = new(ListNode)
		now.Val = (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10

		prev.Next = now
		prev = now
	}
	if carry > 0 {
		now = new(ListNode)
		now.Val = carry
		prev.Next = now
	}
	return root.Next
}

// @lc code=end

// Accepted
// 1568/1568 cases passed (7 ms)
// Your runtime beats 94.75 % of golang submissions
// Your memory usage beats 83.7 % of golang submissions (4.5 MB)
