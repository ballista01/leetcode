/*
 * @lc app=leetcode id=1202 lang=golang
 *
 * [1202] Smallest String With Swaps
 *
 * https://leetcode.com/problems/smallest-string-with-swaps/description/
 *
 * algorithms
 * Medium (52.71%)
 * Likes:    2600
 * Dislikes: 85
 * Total Accepted:    73.4K
 * Total Submissions: 128.7K
 * Testcase Example:  '"dcab"\n[[0,3],[1,2]]'
 *
 * You are given a string s, and an array of pairs of indices in the string
 * pairs where pairs[i] = [a, b] indicates 2 indices(0-indexed) of the string.
 *
 * You can swap the characters at any pair of indices in the given pairs any
 * number of times.
 *
 * Return the lexicographically smallest string that s can be changed to after
 * using the swaps.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "dcab", pairs = [[0,3],[1,2]]
 * Output: "bacd"
 * Explaination:
 * Swap s[0] and s[3], s = "bcad"
 * Swap s[1] and s[2], s = "bacd"
 *
 *
 * Example 2:
 *
 *
 * Input: s = "dcab", pairs = [[0,3],[1,2],[0,2]]
 * Output: "abcd"
 * Explaination:
 * Swap s[0] and s[3], s = "bcad"
 * Swap s[0] and s[2], s = "acbd"
 * Swap s[1] and s[2], s = "abcd"
 *
 * Example 3:
 *
 *
 * Input: s = "cba", pairs = [[0,1],[1,2]]
 * Output: "abc"
 * Explaination:
 * Swap s[0] and s[1], s = "bca"
 * Swap s[1] and s[2], s = "bac"
 * Swap s[0] and s[1], s = "abc"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^5
 * 0 <= pairs.length <= 10^5
 * 0 <= pairs[i][0], pairs[i][1] < s.length
 * s only contains lower case English letters.
 *
 *
 */

package main

import (
	"fmt"
	"sort"
	// "string"
)

func main() {
	s := "qdwyt"
	pairs := [][]int{{2, 3}, {3, 2}, {0, 1}, {4, 0}, {3, 2}}
	res := smallestStringWithSwaps(s, pairs)
	fmt.Println(res)
}

// @lc code=start
func smallestStringWithSwaps(s string, pairs [][]int) string {
	return scanUnions(&pairs, s)
}

func scanUnions(pairs *[][]int, s string) string {
	strlen := len(s)
	var list = [][]int{}
	for index := 0; index < strlen; index++ {
		// list = append(list, make([]int, 1))
		nlist := []int{index}
		list = append(list, nlist)
	}
	for _, pair := range *pairs {
		Union(pair[0], pair[1], &list)
	}
	for index, nlist := range list {
		if nlist[0] != index {
			ancestor := findAncestor(nlist[0], &list)
			list[ancestor] = append(list[ancestor], index)
		}
	}
	res := make([]byte, strlen)
	copy(res, s)
	for _, nlist := range list {
		if len(nlist) > 1 {
			strBytes := make([]byte, 0, strlen)
			for _, bytePos := range nlist {
				strBytes = append(strBytes, s[bytePos])
			}
			sort.Slice(nlist, func(i, j int) bool { return nlist[i] < nlist[j] })
			sort.Slice(strBytes, func(i, j int) bool { return strBytes[i] < strBytes[j] })
			for byteIndex, bytePos := range nlist {
				res[bytePos] = strBytes[byteIndex]
			}
		}
	}
	return string(res)
}

func Union(parent int, child int, list *[][]int) {
	parentsAncestor := findAncestor(parent, list)
	(*list)[findAncestor((*list)[child][0], list)][0] = parentsAncestor
	(*list)[child][0] = parentsAncestor
}

func findAncestor(node int, list *[][]int) int {
	if (*list)[node][0] != node {
		(*list)[node][0] = findAncestor((*list)[node][0], list)
	}
	return (*list)[node][0]
}

// @lc code=end
