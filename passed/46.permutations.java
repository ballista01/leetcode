import java.util.ArrayList;
import java.util.Stack;

/*
 * @lc app=leetcode id=46 lang=java
 *
 * [46] Permutations
 */

// @lc code=start
class Solution {
    public List<List<Integer>> permute(int[] nums) {
        List<List<Integer>> ans = new ArrayList<List<Integer>>();
        if (nums == null || nums.length == 0)
            return ans;
        ArrayList<Integer> perm = new ArrayList<Integer>();
        dfs(nums, perm, ans);
        return ans;
    }

    public void dfs(int[] nums, ArrayList<Integer> perm, List<List<Integer>> ans) {
        if (perm.size() == nums.length) {
            ans.add(new ArrayList<>(perm));
        }
        for (int i = 0; i < nums.length; i++) {
            if (!perm.contains(nums[i])) {
                perm.add(nums[i]);
                dfs(nums, perm, ans);
                perm.remove(perm.size() - 1);
            }
        }
    }
}
// @lc code=end

// Accepted
// 26/26 cases passed (2 ms)
// Your runtime beats 65.04 % of java submissions
// Your memory usage beats 95.52 % of java submissions (42.3 MB)