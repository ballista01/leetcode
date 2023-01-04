
/*
 * @lc app=leetcode id=17 lang=java
 *
 * [17] Letter Combinations of a Phone Number
 */
import java.util.*;

// @lc code=start
class Solution {
    String[][] num2char = {
            { "a", "b", "c" },
            { "d", "e", "f" },
            { "g", "h", "i" },
            { "j", "k", "l" },
            { "m", "n", "o" },
            { "p", "q", "r", "s" },
            { "t", "u", "v" },
            { "w", "x", "y", "z" }
    };

    public List<String> letterCombinations(String digits) {
        List<String> comb = new LinkedList<>();
        for (int i = 0; i < digits.length(); i++) {
            comb = recalculateComb(comb, digits.charAt(i) - '0');
        }
        return comb;
    }

    private List<String> recalculateComb(List<String> oldComb, int num) {
        num -= 2;
        List<String> newComb = new LinkedList<>();

        if (oldComb.isEmpty()) {
            for (String cstr : num2char[num]) {
                newComb.add(cstr);
            }
            return newComb;
        }

        for (String str : oldComb) {
            for (String cstr : num2char[num]) {
                newComb.add(str.concat(cstr));
            }
        }
        return newComb;
    }
}
// @lc code=end

// Accepted
// 25/25 cases passed (1 ms)
// Your runtime beats 95.21 % of java submissions
// Your memory usage beats 95.33 % of java submissions (40.8 MB)