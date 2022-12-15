/*
 * @lc app=leetcode id=200 lang=golang
 *
 * [200] Number of Islands
 */
package main

import (
	"container/list"
)

// func main() {
// 	test := make([][]byte, 4)
// 	test[0] = []byte{1, 1, 0, 0, 0}
// 	test[1] = []byte{1, 1, 0, 0, 0}
// 	test[2] = []byte{0, 0, 1, 0, 0}
// 	test[3] = []byte{0, 0, 0, 1, 1}

// 	numIslands(test)
// }

// @lc code=start
type Coord struct {
	x, y int
}

func numIslands(grid [][]byte) int {
	bytezero := byte('0')
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] > bytezero {
				// fmt.Printf(">>> Found new island at [%d, %d]\n", i, j)
				q := list.New()
				q.PushBack(Coord{i, j})
				grid[i][j] = bytezero
				for q.Len() > 0 {
					top := q.Front().Value.(Coord)
					// fmt.Printf("Mark land [%d, %d] as 0\n", top.x, top.y)
					if top.x-1 >= 0 && grid[top.x-1][top.y] > bytezero {
						q.PushBack(Coord{top.x - 1, top.y})
						grid[top.x-1][top.y] = bytezero
					}
					if top.x+1 < len(grid) && grid[top.x+1][top.y] > bytezero {
						q.PushBack(Coord{top.x + 1, top.y})
						grid[top.x+1][top.y] = bytezero
					}
					if top.y-1 >= 0 && grid[top.x][top.y-1] > bytezero {
						q.PushBack(Coord{top.x, top.y - 1})
						grid[top.x][top.y-1] = bytezero
					}
					if top.y+1 < len(grid[i]) && grid[top.x][top.y+1] > bytezero {
						q.PushBack(Coord{top.x, top.y + 1})
						grid[top.x][top.y+1] = bytezero
					}
					q.Remove(q.Front())
				}
				res++
			}
		}
	}
	return res
}

// @lc code=end

// Accepted
// 49/49 cases passed (7 ms)
// Your runtime beats 62.31 % of golang submissions
// Your memory usage beats 22.74 % of golang submissions (6.6 MB)
