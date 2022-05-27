import java.util.ArrayList;
import java.util.Stack;

/*
 * @lc app=leetcode id=47 lang=java
 *
 * [47] Permutations II
 */

// @lc code=start
class Solution {
    public List<List<Integer>> permuteUnique(int[] nums) {
        List<List<Integer>> ans = new ArrayList<List<Integer>>();
        Stack<Integer> perm = new Stack();
        if (nums == null || nums.length <= 0)
            return ans;
        boolean[] numUsed = new boolean[nums.length];
        Arrays.sort(nums);
        dfs(nums, perm, numUsed, ans);
        return ans;
    }

    public void dfs(int[] nums, Stack<Integer> perm, boolean[] numUsed, List<List<Integer>> ans) {
        if (perm.size() == nums.length) {
            ans.add(new ArrayList<Integer>(perm));
            return;
        }
        for (int i = 0; i < nums.length; i++) {
            if (!numUsed[i]) {
                if (i - 1 >= 0 && nums[i] == nums[i - 1] && !numUsed[i - 1]) {
                    continue;
                } else {
                    perm.push(nums[i]);
                    numUsed[i] = true;
                    dfs(nums, perm, numUsed, ans);
                    perm.pop();
                    numUsed[i] = false;
                }
            }
        }
    }
}
// @lc code=end

// Accepted
// 33/33 cases passed (3 ms)
// Your runtime beats 83.29 % of java submissions
// Your memory usage beats 95.39 % of java submissions (42.7 MB)
