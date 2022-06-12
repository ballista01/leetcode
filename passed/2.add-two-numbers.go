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
	carry, val1, val2 := 0, 0, 0
	dummy := new(ListNode)
	prev := dummy
	for pt1, pt2 := l1, l2; pt1 != nil || pt2 != nil || carry > 0; {
		if pt1 != nil {
			val1 = pt1.Val
		} else {
			val1 = 0
		}
		if pt2 != nil {
			val2 = pt2.Val
		} else {
			val2 = 0
		}

		now := new(ListNode)
		now.Val = (val1 + val2 + carry) % 10
		carry = (val1 + val2 + carry) / 10
		prev.Next = now
		prev = now

		if pt1 != nil {
			pt1 = pt1.Next
		}
		if pt2 != nil {
			pt2 = pt2.Next
		}
	}
	return dummy.Next
}

// @lc code=end

// Accepted
// 1568/1568 cases passed (7 ms)
// Your runtime beats 91.69 % of golang submissions
// Your memory usage beats 73.64 % of golang submissions (4.5 MB)
