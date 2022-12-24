/*
 * @lc app=leetcode id=1351 lang=java
 *
 * [1351] Count Negative Numbers in a Sorted Matrix
 */

// @lc code=start
class Solution {
    public int countNegatives(int[][] grid) {
        int negCount = 0;
        for (int col = 0; col < grid[0].length; col++) {
            int negStart = grid.length;
            for (int row = 0; row < grid.length; row++) {
                if (grid[row][col] < 0) {
                    negStart = row;
                    break;
                }
            }
            negCount += grid.length - negStart;
        }
        return negCount;
    }
}
// @lc code=end

// Accepted
// 44/44 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 86.49 % of java submissions (42.8 MB)