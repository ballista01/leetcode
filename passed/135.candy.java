
/*
 * @lc app=leetcode id=135 lang=java
 *
 * [135] Candy
 */
import java.util.*;

// @lc code=start
class Solution {
    public int candy(int[] ratings) {
        if (ratings.length < 2) {
            return ratings.length;
        }
        int[] candies = new int[ratings.length];
        int totalCandies = 0;
        Arrays.fill(candies, 1);
        for (int i = 1; i < ratings.length; i++) {
            if (ratings[i - 1] < ratings[i]) {
                candies[i] = candies[i - 1] + 1;
            }
        }
        for (int i = ratings.length - 2; i >= 0; i--) {
            if (ratings[i] > ratings[i + 1]) {
                if (!(candies[i] > candies[i + 1])) {
                    candies[i] = candies[i + 1] + 1;
                }
            }
            totalCandies += candies[i + 1];
        }
        totalCandies += candies[0];
        return totalCandies;
    }
}
// @lc code=end

// Accepted
// 48/48 cases passed (2 ms)
// Your runtime beats 99.89 % of java submissions
// Your memory usage beats 82.73 % of java submissions (43.1 MB)
