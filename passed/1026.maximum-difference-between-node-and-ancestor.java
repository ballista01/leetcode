/*
 * @lc app=leetcode id=1026 lang=java
 *
 * [1026] Maximum Difference Between Node and Ancestor
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

// Definition for a binary tree node.

class Solution {
    private int maxDiff = -1;

    public int maxAncestorDiff(TreeNode root) {
        if (root == null) {
            return 0;
        }
        recurHelper(root);
        return maxDiff;
    }

    private int[] recurHelper(TreeNode root) {
        int[] minMax = { root.val, root.val };
        if (root != null && root.left == null && root.right == null) {
            return minMax;
        }
        if (root.left != null) {
            updateMinMaxInPlace(minMax, recurHelper(root.left));
        }
        if (root.right != null) {
            updateMinMaxInPlace(minMax, recurHelper(root.right));
        }
        int diff = getDiff(root.val, minMax);
        if (diff > maxDiff) {
            maxDiff = diff;
        }
        return minMax;
    }

    private void updateMinMaxInPlace(int[] old, int[] another) {
        if (another[0] < old[0]) {
            old[0] = another[0];
        }
        if (another[1] > old[1]) {
            old[1] = another[1];
        }
    }

    private int getDiff(int val, int[] minMax) {
        int diffMin = val - minMax[0];
        if (diffMin < 0)
            diffMin = 0 - diffMin;
        int diffMax = val - minMax[1];
        if (diffMax < 0)
            diffMax = 0 - diffMax;
        return diffMax > diffMin ? diffMax : diffMin;
    }
}
// @lc code=end

// Accepted
// 28/28 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 91.85 % of java submissions (41.9 MB)
