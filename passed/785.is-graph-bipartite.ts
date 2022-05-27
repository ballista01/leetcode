/*
 * @lc app=leetcode id=785 lang=typescript
 *
 * [785] Is Graph Bipartite?
 *
 * https://leetcode.com/problems/is-graph-bipartite/description/
 *
 * algorithms
 * Medium (50.39%)
 * Likes:    4592
 * Dislikes: 271
 * Total Accepted:    305.7K
 * Total Submissions: 593.2K
 * Testcase Example:  '[[1,2,3],[0,2],[0,1,3],[0,2]]'
 *
 * There is an undirected graph with n nodes, where each node is numbered
 * between 0 and n - 1. You are given a 2D array graph, where graph[u] is an
 * array of nodes that node u is adjacent to. More formally, for each v in
 * graph[u], there is an undirected edge between node u and node v. The graph
 * has the following properties:
 *
 *
 * There are no self-edges (graph[u] does not contain u).
 * There are no parallel edges (graph[u] does not contain duplicate
 * values).
 * If v is in graph[u], then u is in graph[v] (the graph is undirected).
 * The graph may not be connected, meaning there may be two nodes u and v such
 * that there is no path between them.
 *
 *
 * A graph is bipartite if the nodes can be partitioned into two independent
 * sets A and B such that every edge in the graph connects a node in set A and
 * a node in set B.
 *
 * Return true if and only if it is bipartite.
 *
 *
 * Example 1:
 *
 *
 * Input: graph = [[1,2,3],[0,2],[0,1,3],[0,2]]
 * Output: false
 * Explanation: There is no way to partition the nodes into two independent
 * sets such that every edge connects a node in one and a node in the other.
 *
 * Example 2:
 *
 *
 * Input: graph = [[1,3],[0,2],[1,3],[0,2]]
 * Output: true
 * Explanation: We can partition the nodes into two sets: {0, 2} and {1, 3}.
 *
 *
 * Constraints:
 *
 *
 * graph.length == n
 * 1 <= n <= 100
 * 0 <= graph[u].length < n
 * 0 <= graph[u][i] <= n - 1
 * graph[u] does not contain u.
 * All the values of graph[u] are unique.
 * If graph[u] contains v, then graph[v] contains u.
 *
 *
 */

const g = [
  [1, 3],
  [0, 2],
  [1, 3],
  [0, 2],
];
console.log(isBipartite(g));
console.log("terminated...");

// @lc code=start
function isBipartite(graph: number[][]): boolean {
  const color = Array<number>();

  for (let i = 0; i < graph.length; i++) {
    if (color[i] != undefined) continue;
    const queue = Array<number>();
    queue.push(i);
    color[i] = 1;
    while (queue.length > 0) {
      const u: number | undefined = queue.shift();
      if (u === undefined) throw new Error("BFS queue popped undefined!");
      for (const v of graph[u]) {
        if (color[v] != undefined) {
          // v is colored
          if (color[v] == color[u]) {
            // u and v have different colors
            return false;
          }
        } else {
          // v is not colored, give it different color to u
          color[v] = 0 - color[u];
          queue.push(v);
        }
      }
    }
  }
  return true;
}
// @lc code=end

// Accepted
// 80/80 cases passed (73 ms)
// Your runtime beats 96.15 % of typescript submissions
// Your memory usage beats 73.08 % of typescript submissions (45.6 MB)
