/*
 * @lc app=leetcode id=101 lang=java
 *
 * [101] Symmetric Tree
 */
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
    public boolean isSymmetric(TreeNode root) {
        return symmetricHelper(root.left, root.right);
    }

    public boolean symmetricHelper(TreeNode a, TreeNode b) {
        if (a == null && b == null)
            return true;
        if (a == null || b == null)
            return false;
        if (a.val != b.val)
            return false;
        return symmetricHelper(a.left, b.right) && symmetricHelper(a.right, b.left);
    }
}
// @lc code=end

// Accepted
// 198/198 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 95.06 % of java submissions (40.4 MB)