
/*
 * @lc app=leetcode id=1382 lang=java
 *
 * [1382] Balance a Binary Search Tree
 */
import java.util.*;

class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode() {
    }

    TreeNode(int val) {
        this.val = val;
    }

    TreeNode(int val, TreeNode left, TreeNode right) {
        this.val = val;
        this.left = left;
        this.right = right;
    }
}
// @lc code=start

class Solution {
    private List<Integer> sorted;

    public TreeNode balanceBST(TreeNode root) {
        sorted = new ArrayList<>();
        inOrderTraverse(root);
        TreeNode newRoot = sortedListToBST(0, sorted.size() - 1);
        return newRoot;
    }

    private void inOrderTraverse(TreeNode root) {
        if (root == null) {
            return;
        }
        inOrderTraverse(root.left);
        sorted.add(root.val);
        inOrderTraverse(root.right);
    }

    private TreeNode sortedListToBST(int lo, int hi) {
        if (lo > hi) {
            return null;
        }
        int mid = (lo + hi + 1) / 2;
        TreeNode root = new TreeNode(sorted.get(mid));
        root.left = sortedListToBST(lo, mid - 1);
        root.right = sortedListToBST(mid + 1, hi);
        return root;
    }
}
// @lc code=end

// Accepted
// 17/17 cases passed (3 ms)
// Your runtime beats 86.73 % of java submissions
// Your memory usage beats 80.81 % of java submissions (45.8 MB)