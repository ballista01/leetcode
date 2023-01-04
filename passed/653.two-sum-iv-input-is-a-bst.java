import java.util.*;

/*
 * @lc app=leetcode id=653 lang=java
 *
 * [653] Two Sum IV - Input is a BST
 */
public class TreeNode {
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

    public boolean findTarget(TreeNode root, int k) {
        if (root == null || root.left == null && root.right == null) {
            return false;
        }
        sorted = new ArrayList<>();
        inOrderTraversal(root);

        int lo = 0;
        int hi = sorted.size() - 1;
        int thisSum;
        while (lo < hi) {
            thisSum = sorted.get(lo) + sorted.get(hi);
            if (thisSum == k) {
                return true;
            }
            if (thisSum < k) {
                lo++;
            } else {
                hi--;
            }
        }
        return false;
    }

    private void inOrderTraversal(TreeNode root) {
        if (root == null) {
            return;
        }
        inOrderTraversal(root.left);
        sorted.add(root.val);
        inOrderTraversal(root.right);
    }
}
// @lc code=end

// Accepted
// 423/423 cases passed (3 ms)
// Your runtime beats 96.63 % of java submissions
// Your memory usage beats 81.92 % of java submissions (43.1 MB)