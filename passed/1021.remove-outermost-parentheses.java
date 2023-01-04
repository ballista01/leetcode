
/*
 * @lc app=leetcode id=1021 lang=java
 *
 * [1021] Remove Outermost Parentheses
 */
import java.util.*;

// @lc code=start
class Solution {
    public String removeOuterParentheses(String s) {
        LinkedList<Character> pList = new LinkedList<>();
        int bal = 0;
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < s.length(); i++) {
            char ch = s.charAt(i);
            if (ch == '(') {
                bal--;
            } else {
                bal++;
            }
            pList.addLast(ch);
            if (bal == 0) {
                pList.removeFirst();
                pList.removeLast();
                for (char c : pList) {
                    sb.append(c);
                }
                pList.clear();
            }
        }

        return sb.toString();
    }
}
// @lc code=end

// Accepted
// 59/59 cases passed (6 ms)
// Your runtime beats 74.71 % of java submissions
// Your memory usage beats 90.23 % of java submissions (42 MB)