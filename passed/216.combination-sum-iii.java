
/*
 * @lc app=leetcode id=216 lang=java
 *
 * [216] Combination Sum III
 *
 * https://leetcode.com/problems/combination-sum-iii/description/
 *
 * algorithms
 * Medium (64.02%)
 * Likes:    3100
 * Dislikes: 76
 * Total Accepted:    302.5K
 * Total Submissions: 466.9K
 * Testcase Example:  '3\n7'
 *
 * Find all valid combinations of k numbers that sum up to n such that the
 * following conditions are true:
 * 
 * 
 * Only numbers 1 through 9 are used.
 * Each number is used at most once.
 * 
 * 
 * Return a list of all possible valid combinations. The list must not contain
 * the same combination twice, and the combinations may be returned in any
 * order.
 * 
 * 
 * Example 1:
 * 
 * 
 * Input: k = 3, n = 7
 * Output: [[1,2,4]]
 * Explanation:
 * 1 + 2 + 4 = 7
 * There are no other valid combinations.
 * 
 * Example 2:
 * 
 * 
 * Input: k = 3, n = 9
 * Output: [[1,2,6],[1,3,5],[2,3,4]]
 * Explanation:
 * 1 + 2 + 6 = 9
 * 1 + 3 + 5 = 9
 * 2 + 3 + 4 = 9
 * There are no other valid combinations.
 * 
 * 
 * Example 3:
 * 
 * 
 * Input: k = 4, n = 1
 * Output: []
 * Explanation: There are no valid combinations.
 * Using 4 different numbers in the range [1,9], the smallest sum we can get is
 * 1+2+3+4 = 10 and since 10 > 1, there are no valid combination.
 * 
 * 
 * 
 * Constraints:
 * 
 * 
 * 2 <= k <= 9
 * 1 <= n <= 60
 * 
 * 
 */
import java.util.List;
import java.util.ArrayList;
import java.util.Stack;
import java.lang.Integer;
// @lc code=start

class Solution {
    public List<List<Integer>> combinationSum3(int k, int n) {
        List<List<Integer>> res = new ArrayList<>();
        Stack<Integer> track = new Stack<Integer>();
        trackingHelper(k, n, 9, track, res);
        return res;
    }

    public void trackingHelper(int k, int n, int maxNum, Stack<Integer> track, List<List<Integer>> res) {
        if (k < 0 || n < 0)
            return;
        if (k == 0 && n == 0) {
            res.add(new ArrayList<Integer>(track));
        }
        for (int i = 1; i <= maxNum; i++) {
            track.push(i);
            trackingHelper(k - 1, n - i, i - 1, track, res);
            track.pop();
        }
    }

}
// @lc code=end

// Accepted
// 18/18 cases passed (1 ms)
// Your runtime beats 73.61 % of java submissions
// Your memory usage beats 71.91 % of java submissions (40.9 MB)