/*
 * @lc app=leetcode id=1631 lang=golang
 *
 * [1631] Path With Minimum Effort
 *
 * https://leetcode.com/problems/path-with-minimum-effort/description/
 *
 * algorithms
 * Medium (51.95%)
 * Likes:    2882
 * Dislikes: 119
 * Total Accepted:    97.2K
 * Total Submissions: 178K
 * Testcase Example:  '[[1,2,2],[3,8,2],[5,3,5]]'
 *
 * You are a hiker preparing for an upcoming hike. You are given heights, a 2D
 * array of size rows x columns, where heights[row][col] represents the height
 * of cell (row, col). You are situated in the top-left cell, (0, 0), and you
 * hope to travel to the bottom-right cell, (rows-1, columns-1) (i.e.,
 * 0-indexed). You can move up, down, left, or right, and you wish to find a
 * route that requires the minimum effort.
 *
 * A route's effort is the maximum absolute difference in heights between two
 * consecutive cells of the route.
 *
 * Return the minimum effort required to travel from the top-left cell to the
 * bottom-right cell.
 *
 *
 * Example 1:
 *
 *
 *
 *
 * Input: heights = [[1,2,2],[3,8,2],[5,3,5]]
 * Output: 2
 * Explanation: The route of [1,3,5,3,5] has a maximum absolute difference of 2
 * in consecutive cells.
 * This is better than the route of [1,2,2,2,5], where the maximum absolute
 * difference is 3.
 *
 *
 * Example 2:
 *
 *
 *
 *
 * Input: heights = [[1,2,3],[3,8,4],[5,3,5]]
 * Output: 1
 * Explanation: The route of [1,2,3,4,5] has a maximum absolute difference of 1
 * in consecutive cells, which is better than route [1,3,5,3,5].
 *
 *
 * Example 3:
 *
 *
 * Input: heights =
 * [[1,2,1,1,1],[1,2,1,2,1],[1,2,1,2,1],[1,2,1,2,1],[1,1,1,2,1]]
 * Output: 0
 * Explanation: This route does not require any effort.
 *
 *
 *
 * Constraints:
 *
 *
 * rows == heights.length
 * columns == heights[i].length
 * 1 <= rows, columns <= 100
 * 1 <= heights[i][j] <= 10^6
 *
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	heights := [][]int{{1, 10, 6, 7, 9, 10, 4, 9}}
	fmt.Println(minimumEffortPath(heights))
}

// @lc code=start
func minimumEffortPath(heights [][]int) int {
	if len(heights) == 0 {
		return -1
	}
	minEffort := dijkstra([2]int{0, 0}, [2]int{len(heights) - 1, len(heights[0]) - 1}, heights)
	return minEffort
}

func dijkstra(source [2]int, dest [2]int, heights [][]int) int {
	inf := math.MaxInt
	dirMtx := [4][2]int{
		{-1, 0},
		{0, -1}, {0, 1},
		{1, 0},
	}
	rows, cols := len(heights), len(heights[0])

	dist := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = inf
		}
	}

	dist[source[0]][source[1]] = 0
	queue := [][2]int{{0, 0}}

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		currDiff := dist[u[0]][u[1]]

		for _, dir := range dirMtx {
			i := u[0] + dir[0]
			j := u[1] + dir[1]
			if i >= 0 && i < rows && j >= 0 && j < cols {
				cellAbsDiff := heights[u[0]][u[1]] - heights[i][j]
				if cellAbsDiff < 0 {
					cellAbsDiff = 0 - cellAbsDiff
				}
				newDiff := currDiff
				if cellAbsDiff > currDiff {
					newDiff = cellAbsDiff
				}
				if newDiff < dist[i][j] {
					dist[i][j] = newDiff
					queue = append(queue, [2]int{i, j})
				}
			}

		}
	}

	return dist[dest[0]][dest[1]]
}

// @lc code=end

// Success
// Details
// Runtime: 95 ms, faster than 94.74% of Go online submissions for Path With Minimum Effort.
// Memory Usage: 8.4 MB, less than 39.47% of Go online submissions for Path With Minimum Effort.
