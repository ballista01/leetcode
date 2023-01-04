
/*
 * @lc app=leetcode id=42 lang=java
 *
 * [42] Trapping Rain Water
 */
import java.util.*;

// @lc code=start
class Solution {
    Set<Integer> peakSet;
    Map<Integer, Integer> heightCount;
    List<Integer> peaks;

    public int trap(int[] height) {
        int vol = 0;
        peaks = getPeaks(height);
        for (int i = 1; i < peaks.size(); i++) {
            int minPeakHeight = height[peaks.get(i)];
            if (height[peaks.get(i - 1)] < minPeakHeight) {
                minPeakHeight = peaks.get(i - 1);
            }

            for (int j = peaks.get(i - 1) + 1; j < peaks.get(i) - 1; j++) {
                vol += minPeakHeight - height[j];
            }
        }
    }

    private List<Integer> getPeaks(int[] height) {
        // int leftHeight, thisHeight, rightHeight;
        // List<Integer> res = new ArrayList<>();
        // for (int i = 1; i < height.length - 1; i++) {
        // leftHeight = height[i - 1];
        // thisHeight = height[i];
        // rightHeight = height[i + 1];

        // if (thisHeight >= leftHeight && thisHeight >= rightHeight
        // && leftHeight != rightHeight) {
        // res.add(i);
        // }
        // }
        // return res;

        peakSet = new TreeSet<>();
        heightCount = new HashMap<>();
        for (int h : height) {
            Integer count = heightCount.get(h);
            if (count == null) {
                heightCount.put(h, 1);
            } else {
                heightCount.put(h, count + 1);
            }
        }

        heightCount.clone();
    }

}
// @lc code=end
