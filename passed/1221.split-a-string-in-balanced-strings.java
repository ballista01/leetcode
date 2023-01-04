/*
 * @lc app=leetcode id=1221 lang=java
 *
 * [1221] Split a String in Balanced Strings
 */

// @lc code=start
class Solution {
    public int balancedStringSplit(String s) {
        int bal = 0;
        int numSubStr = 0;
        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            if (ch == 'L') {
                bal--;
            } else {
                bal++;
            }
            if (bal == 0) {
                numSubStr++;
            }
        }
        return numSubStr;
    }
}
// @lc code=end

// Accepted
// 40/40 cases passed (1 ms)
// Your runtime beats 56.23 % of java submissions
// Your memory usage beats 52.14 % of java submissions (41.5 MB)