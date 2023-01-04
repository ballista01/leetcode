/*
 * @lc app=leetcode id=459 lang=java
 *
 * [459] Repeated Substring Pattern
 */

// @lc code=start
class Solution {

    public boolean repeatedSubstringPattern(String s) {
        if (s.length() == 1)
            return false;

        for (int repetition = 2; repetition <= s.length(); repetition++) {
            if (s.length() % repetition != 0) {
                continue;
            }
            int subStrLen = s.length() / repetition;
            String subStr = s.substring(0, subStrLen);
            boolean matched = true;
            // System.out.println("checking: " + subStr);
            for (int i = subStrLen; i < s.length(); i += subStrLen) {
                // System.out.printf("i=%d, end=%d\n", i, i + subStrLen);
                if (!subStr.equals(s.substring(i, i + subStrLen))) {
                    matched = false;
                    break;
                }
            }
            if (matched) {
                return true;
            }
        }
        return false;
    }
}
// @lc code=end

// Accepted
// 129/129 cases passed (25 ms)
// Your runtime beats 74.76 % of java submissions
// Your memory usage beats 25.9 % of java submissions (54.3 MB)