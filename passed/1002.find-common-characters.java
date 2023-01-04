
/*
 * @lc app=leetcode id=1002 lang=java
 *
 * [1002] Find Common Characters
 */
import java.util.*;

// @lc code=start
class Solution {
    public List<String> commonChars(String[] words) {
        int wHash[][] = new int[words.length][26];

        for (int i = 0; i < words.length; i++) {
            wHash[i] = new int[26];
            for (int j = 0; j < words[i].length(); j++) {
                wHash[i][words[i].charAt(j) - 'a']++;
            }
        }

        List<String> resList = new ArrayList<>();
        for (int i = 0; i < 26; i++) {
            char thisChar = (char) ('a' + i);
            int minOccur = Integer.MAX_VALUE;
            for (int j = 0; j < wHash.length; j++) {
                if (wHash[j][i] < minOccur) {
                    minOccur = wHash[j][i];
                }
            }
            for (int j = 0; j < minOccur; j++) {
                resList.add("" + thisChar);
            }
        }
        return resList;
    }
}
// @lc code=end

// Accepted
// 83/83 cases passed (4 ms)
// Your runtime beats 79.13 % of java submissions
// Your memory usage beats 74.19 % of java submissions (42.6 MB)