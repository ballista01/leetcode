/*
 * @lc app=leetcode id=1642 lang=golang
 *
 * [1642] Furthest Building You Can Reach
 */
package main

// @lc code=start

type Cost struct {
	priority int
	index    int
}

type PriorityQueue []*Cost

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	cost := x.(*Cost)
	cost.index = n
	*pq = append(*pq, cost)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	cost := old[n-1]
	old[n-1] = nil
	cost.index = -1
	*pq = old[0 : n-1]
	return cost
}

func furthestBuilding(heights []int, bricks int, ladders int) int {

}

// @lc code=end
