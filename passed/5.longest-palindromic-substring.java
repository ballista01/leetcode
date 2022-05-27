// @lc code=start
class Solution {
    public String longestPalindrome(String s) {
        if (s.length() == 0)
            return "";
        int[] longestPalIndex = new int[] { 0, 1 };
        int[] palIndex = new int[] { 0, 1 };
        for (int i = 1; i < s.length(); i++) {
            palIndex[0] = i;
            palIndex[1] = i + 1;
            // odd palindrome check
            for (int numLeft = 1; i - numLeft >= 0 && i + numLeft < s.length()
                    && s.charAt(i - numLeft) == s.charAt(i + numLeft); numLeft++) {
                // mirrored edge, increase palindrome size
                palIndex[0]--;
                palIndex[1]++;
                // compare to the longest and update
                if (palIndex[1] - palIndex[0] > longestPalIndex[1] - longestPalIndex[0]) {
                    // longestPal = nowPal;
                    longestPalIndex[0] = palIndex[0];
                    longestPalIndex[1] = palIndex[1];
                }
            }

            palIndex[0] = i;
            palIndex[1] = i;
            // even palindrome check
            for (int numLeft = 1; i - numLeft >= 0 && i + numLeft - 1 < s.length()
                    && s.charAt(i - numLeft) == s.charAt(i + numLeft - 1); numLeft++) {
                palIndex[0]--;
                palIndex[1]++;
                if (palIndex[1] - palIndex[0] > longestPalIndex[1] - longestPalIndex[0]) {
                    longestPalIndex[0] = palIndex[0];
                    longestPalIndex[1] = palIndex[1];
                }
            }
        }

        return s.substring(longestPalIndex[0], longestPalIndex[1]);
    }
}
// @lc code=end

// Accepted
// 180/180 cases passed (54 ms)
// Your runtime beats 52.78 % of java submissions
// Your memory usage beats 77.75 % of java submissions (42.8 MB)