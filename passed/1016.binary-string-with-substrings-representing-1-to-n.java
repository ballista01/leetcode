/*
 * @lc app=leetcode id=1016 lang=java
 *
 * [1016] Binary String With Substrings Representing 1 To N
 */

// @lc code=start
class Solution {
    public boolean queryString(String s, int n) {
        for (int i = n; i > 0; i--) {
            if (!isSubstring(s, Integer.toBinaryString(i))) {
                return false;
            }
        }
        return true;
    }

    private boolean isSubstring(String complete, String part) {
        if (part.length() > complete.length())
            return false;
        int offset = complete.length() - part.length();
        for (int i = 0; i <= offset; i++) {
            if (complete.substring(i, i + part.length()).equals(part)) {
                return true;
            }
        }
        return false;
    }
}
// @lc code=end

// Accepted
// 26/26 cases passed (1 ms)
// Your runtime beats 82.61 % of java submissions
// Your memory usage beats 80.19 % of java submissions (40.7 MB)