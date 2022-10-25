/*
 * @lc app=leetcode id=121 lang=cpp
 *
 * [121] Best Time to Buy and Sell Stock
 */
#include <iostream>
#include <vector>
#include <math.h>

using namespace std;
// @lc code=start
class Solution
{
public:
    int maxProfit(vector<int> &prices)
    {
        int minPrice = INT32_MAX;
        int maxProfit = 0;
        for (int v : prices)
        {
            if (v < minPrice)
            {
                minPrice = v;
            }
            if (v - minPrice > maxProfit)
            {
                maxProfit = v - minPrice;
            }
        }
        return maxProfit;
    }
};
// @lc code=end

// Accepted
// 211/211 cases passed (280 ms)
// Your runtime beats 40.09 % of cpp submissions
// Your memory usage beats 51.09 % of cpp submissions (93.3 MB)