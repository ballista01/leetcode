/*
 * @lc app=leetcode id=1025 lang=java
 *
 * [1025] Divisor Game
 */

// @lc code=start
class Solution {
    public boolean divisorGame(int n) {
        // if(n==1) return false;
        boolean aliceMovedLast = false;
        while (true) {
            int x = 1;
            boolean xFound = false;
            for (; x < n; x++) {
                if (n % x == 0) {
                    // System.out.printf("%s choose %d\n", aliceMovedLast ? "Bob" : "Alice", x);
                    n = n - x;
                    aliceMovedLast = !aliceMovedLast;
                    xFound = true;
                    break;
                }
            }
            if (!xFound) {
                return aliceMovedLast;
            }
        }
    }
}
// @lc code=end

// Accepted
// 40/40 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 53.67 % of java submissions (40.5 MB)