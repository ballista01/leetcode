/*
 * @lc app=leetcode id=59 lang=java
 *
 * [59] Spiral Matrix II
 */

// @lc code=start
class Solution {
    int[][] dir = { { 0, 1 }, { 1, 0 }, { 0, -1 }, { -1, 0 } };
    int dirIndex = 0;

    public int[][] generateMatrix(int n) {
        int total = n * n;
        int[][] matrix = new int[n][n];
        int[] coord = { 0, 0 };
        int nextX;
        int nextY;
        for (int i = 1; i <= total; i++) {
            matrix[coord[0]][coord[1]] = i;
            nextX = coord[0] + dir[dirIndex][0];
            nextY = coord[1] + dir[dirIndex][1];
            if (nextX < 0 || nextX >= n || nextY < 0 || nextY >= n || matrix[nextX][nextY] != 0) {
                dirIndex++;
                dirIndex %= dir.length;
                coord[0] += dir[dirIndex][0];
                coord[1] += dir[dirIndex][1];
            } else {
                coord[0] = nextX;
                coord[1] = nextY;
            }
        }
        return matrix;
    }
}
// @lc code=end

// Accepted
// 20/20 cases passed (0 ms)
// Your runtime beats 100 % of java submissions
// Your memory usage beats 71.19 % of java submissions (40.4 MB)