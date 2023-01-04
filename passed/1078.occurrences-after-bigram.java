
/*
 * @lc app=leetcode id=1078 lang=java
 *
 * [1078] Occurrences After Bigram
 */
import java.util.*;

// @lc code=start
class Solution {

    public String[] findOcurrences(String text, String first, String second) {
        List<String> resList = new LinkedList<>();
        String[] words = text.split(" ");
        for (int i = 1; i < words.length - 1; i++) {
            if (words[i].equals(second) && words[i - 1].equals(first)) {
                resList.add(words[i + 1]);
            }
        }
        return resList.stream().toArray(String[]::new);
    }
}
// @lc code=end

// Accepted
// 30/30 cases passed (2 ms)
// Your runtime beats 20.69 % of java submissions
// Your memory usage beats 77.74 % of java submissions (40.7 MB)