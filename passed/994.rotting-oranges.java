
/*
 * @lc app=leetcode id=994 lang=java
 *
 * [994] Rotting Oranges
 */
import java.util.*;

// @lc code=start
class Solution {
    int[][] dir = { { 0, 1 }, { 1, 0 }, { 0, -1 }, { -1, 0 } };
    int totalOranges = 0;
    int rottenOranges = 0;
    int numMinutes = 0;

    public int orangesRotting(int[][] grid) {

        LinkedList<Coord> q = new LinkedList<>();
        for (int i = 0; i < grid.length; i++) {
            for (int j = 0; j < grid[i].length; j++) {
                if (grid[i][j] > 0) {
                    totalOranges++;
                    if (grid[i][j] == 2) {
                        q.addLast(new Coord(i, j, 0));
                        rottenOranges++;
                    }
                }
            }
        }
        while (!q.isEmpty()) {
            Coord oc = q.removeFirst();
            probeAndEnqueue(grid, oc, q);
        }
        return totalOranges == rottenOranges ? numMinutes : -1;
    }

    private void probeAndEnqueue(int[][] grid, Coord oc, LinkedList<Coord> q) {
        for (int i = 0; i < dir.length; i++) {
            Coord newCoord = new Coord(oc.x + dir[i][0], oc.y + dir[i][1], oc.layer + 1);
            if (newCoord.x < 0 || newCoord.x >= grid.length || newCoord.y < 0 || newCoord.y >= grid[newCoord.x].length)
                continue;
            if (grid[newCoord.x][newCoord.y] == 1) {
                q.addLast(newCoord);
                grid[newCoord.x][newCoord.y] = 2;
                rottenOranges++;
                numMinutes = newCoord.layer;
            }
        }
    }

}

class Coord {
    public Coord(int x, int y, int layer) {
        this.x = x;
        this.y = y;
        this.layer = layer;
    }

    public int x, y, layer;
}
// @lc code=end

// Accepted
// 306/306 cases passed (4 ms)
// Your runtime beats 44.32 % of java submissions
// Your memory usage beats 27.99 % of java submissions (43.2 MB)