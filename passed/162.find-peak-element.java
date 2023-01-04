/*
 * @lc app=leetcode id=162 lang=java
 *
 * [162] Find Peak Element
 */

// @lc code=start
class Solution {
    public int findPeakElement(int[] nums) {
        int frontVal, thisVal, backVal;
        int res = -1;
        for (int i = 1; i < nums.length - 1; i++) {
            frontVal = nums[i - 1];
            thisVal = nums[i];
            backVal = nums[i + 1];
            if (thisVal > frontVal && thisVal > backVal) {
                return i;
            }
        }
        if (res == -1) {
            frontVal = nums[0];
            backVal = nums[nums.length - 1];
            if (frontVal > backVal) {
                res = 0;
            } else {
                res = nums.length - 1;
            }
        }
        return res;
    }
}
// @lc code=end

// Accepted
// 65/65 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 84.3 % of java submissions (41.7 MB)