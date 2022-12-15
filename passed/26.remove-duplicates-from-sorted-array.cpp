/*
 * @lc app=leetcode id=26 lang=cpp
 *
 * [26] Remove Duplicates from Sorted Array
 */
#include <vector>
using namespace std;
// @lc code=start
class Solution
{
public:
    int removeDuplicates(vector<int> &nums)
    {
        int prev = -111, distinctIndex = 0;
        for (int i = 0; i < nums.size(); i++)
        {
            if (nums[i] != prev)
            {
                prev = nums[i];
                nums[distinctIndex++] = nums[i];
            }
        }
        return distinctIndex;
    }
};
// @lc code=end

// Accepted
// 361/361 cases passed (13 ms)
// Your runtime beats 86.73 % of cpp submissions
// Your memory usage beats 99.8 % of cpp submissions (18.2 MB)