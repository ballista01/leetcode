import java.util.ArrayList;
import java.util.LinkedList;
import java.util.List;

/*
 * @lc app=leetcode id=93 lang=java
 *
 * [93] Restore IP Addresses
 */

// @lc code=start
class Solution {
    private int[] startIndices;
    private List<int[]> resList;
    private List<String> outputList;

    public List<String> restoreIpAddresses(String s) {
        if (s.length() > 12) {
            return new ArrayList<String>();
        }
        startIndices = new int[4];
        startIndices[0] = 0;
        resList = new LinkedList<>();
        recurSolve(s, 1);
        outputList = new ArrayList<>(resList.size());
        for (int[] arr : resList) {
            String str = String.format("%s.%s.%s.%s",
                    s.substring(arr[0], arr[1]),
                    s.substring(arr[1], arr[2]),
                    s.substring(arr[2], arr[3]),
                    s.substring(arr[3]));
            outputList.add(str);
        }
        return outputList;
    }

    private void recurSolve(String s, int siIndex) {
        int minSI = startIndices[siIndex - 1] + 1;
        int maxSI = startIndices[siIndex - 1] + 3;
        if (minSI >= s.length()) {
            return;
        }
        if (maxSI >= s.length()) {
            maxSI = s.length() - 1;
        }
        if (maxSI < minSI) {
            return;
        }

        for (int i = minSI; i <= maxSI; i++) {
            String thisPart = s.substring(startIndices[siIndex - 1], i);
            boolean isValid = isValidPart(thisPart);
            if (siIndex == 3) {
                isValid = isValid && isValidPart(s.substring(i));
            }
            if (isValid) {
                if (siIndex == 3) {
                    startIndices[siIndex] = i;
                    resList.add(startIndices.clone());
                } else {
                    startIndices[siIndex] = i;
                    recurSolve(s, siIndex + 1);
                }
            }
        }
    }

    private boolean isValidPart(String s) {
        if (s.length() == 0 && s.length() > 3) {
            return false;
        }
        if (s.length() > 1 && s.charAt(0) == '0') {
            return false;
        }

        int val = Integer.parseInt(s);
        if (val < 0 || val > 255) {
            return false;
        }

        return true;
    }

}
// @lc code=end

// Accepted
// 145/145 cases passed (6 ms)
// Your runtime beats 62.45 % of java submissions
// Your memory usage beats 86.74 % of java submissions (42.4 MB)