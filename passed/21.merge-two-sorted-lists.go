/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	root := new(ListNode)
	prev := root
	var now *ListNode
	now1, now2 := list1, list2
	for now1 != nil && now2 != nil {
		now := new(ListNode)
		if now1.Val < now2.Val {
			now.Val = now1.Val
			now1 = now1.Next
		} else {
			now.Val = now2.Val
			now2 = now2.Next
		}
		prev.Next = now
		prev = prev.Next
	}
	for ; now1 != nil; now1 = now1.Next {
		now = new(ListNode)
		now.Val = now1.Val
		prev.Next = now
		prev = prev.Next
	}
	for ; now2 != nil; now2 = now2.Next {
		now = new(ListNode)
		now.Val = now2.Val
		prev.Next = now
		prev = prev.Next
	}
	return root.Next
}

// @lc code=end

// Accepted
// 208/208 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 49.71 % of golang submissions (2.5 MB)
