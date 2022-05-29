/*
 * @lc app=leetcode id=1559 lang=golang
 *
 * [1559] Detect Cycles in 2D Grid
 */
package main

import "container/list"

// @lc code=start
type Vertex struct {
	x, y int
}

func containsCycle(grid [][]byte) bool {
	visited := make([][]bool, len(grid))
	for i := range visited {
		visited[i] = make([]bool, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[0] {
			if !visited[i][j] {
				if BFS(Vertex{i, j}, grid, visited) {
					return true
				}
			}
		}
	}
	return false
}

func BFS(start Vertex, grid [][]byte, visited [][]bool) bool {
	q := list.New()
	q.PushBack(start)
	for q.Len() != 0 {
		v := q.Front().Value.(Vertex)
		q.Remove(q.Front())
		if visited[v.x][v.y] {
			return true
		}
		visited[v.x][v.y] = true

		for dirIndex := 0; dirIndex < 4; dirIndex++ {
			i, j := v.x, v.y
			switch dirIndex {
			case 0:
				j++
			case 1:
				i++
			case 2:
				j--
			case 3:
				i--
			}
			if i >= 0 && i < len(visited) && j >= 0 && j < len(visited[0]) && !visited[i][j] && grid[v.x][v.y] == grid[i][j] {
				q.PushBack(Vertex{i, j})
			}
		}
	}
	return false
}

// @lc code=end

// Accepted
// 76/76 cases passed (70 ms)
// Your runtime beats 58.82 % of golang submissions
// Your memory usage beats 94.12 % of golang submissions (9.2 MB)
