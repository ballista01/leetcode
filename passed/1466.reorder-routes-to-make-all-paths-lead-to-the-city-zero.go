/*
 * @lc app=leetcode id=1466 lang=golang
 *
 * [1466] Reorder Routes to Make All Paths Lead to the City Zero
 */
package main

import "container/list"

func main() {
	minReorder(6, [][]int{{0, 1}, {1, 3}, {2, 3}, {4, 0}, {4, 5}})
}

// @lc code=start
func minReorder(n int, connections [][]int) int {
	numReverse := 0
	adjList := make([][]int, n)
	for _, v := range connections {
		adjList[v[0]] = append(adjList[v[0]], -v[1])
		adjList[v[1]] = append(adjList[v[1]], v[0])
	}

	q := list.New()
	q.PushBack(0)
	visited := make([]bool, n)
	for q.Len() != 0 {
		nowCity := q.Remove(q.Front()).(int)
		for _, v := range adjList[nowCity] {
			if v < 0 {
				v = 0 - v
				if !visited[v] {
					numReverse++
				}
			}
			if visited[v] {
				continue
			}
			q.PushBack(v)
			visited[v] = true
		}
	}
	return numReverse
}

// @lc code=end

// Accepted
// 76/76 cases passed (213 ms)
// Your runtime beats 76.67 % of golang submissions
// Your memory usage beats 90 % of golang submissions (13.4 MB)
