/*
 * @lc app=leetcode id=94 lang=golang
 *
 * [94] Binary Tree Inorder Traversal
 */
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// @lc code=start
func inorderTraversal(root *TreeNode) []int {
	seq := make([]int, 0, 101)
	traverse(root, &seq)
	return seq
}
func traverse(root *TreeNode, seq *[]int) {
	if root == nil {
		return
	}
	traverse(root.Left, seq)
	*seq = append(*seq, root.Val)
	traverse(root.Right, seq)
}

// @lc code=end

// Accepted
// 70/70 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 28.52 % of golang submissions (2 MB)
