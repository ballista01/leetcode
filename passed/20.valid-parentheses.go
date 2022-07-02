/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */
package pending

import "container/list"

// @lc code=start
func isValid(s string) bool {
	st := list.New()
	var pmap = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, v := range s {
		if v == '(' || v == '[' || v == '{' {
			st.PushBack(v)
		} else {
			if st.Len() > 0 && st.Back().Value.(rune) == pmap[v] {
				st.Remove(st.Back())
			} else {
				return false
			}
		}
	}

	return st.Len() == 0
}

// Accepted
// 92/92 cases passed (0 ms)
// Your runtime beats 100 % of golang submissions
// Your memory usage beats 47.42 % of golang submissions (2.1 MB)

// @lc code=end

// tested with switch case, no significant performance difference

// func isValid(s string) bool {
// 	st := list.New()
// 	for _, v := range s {
// 		if v == '(' || v == '[' || v == '{' {
// 			st.PushBack(v)
// 		} else {
// 			if st.Len() > 0 {
// 				switch st.Back().Value.(rune) {
// 				case '(':
// 					if v != ')' {
// 						return false
// 					}
// 				case '[':
// 					if v != ']' {
// 						return false
// 					}
// 				case '{':
// 					if v != '}' {
// 						return false
// 					}
// 				default:
// 					return false
// 				}
// 				st.Remove(st.Back())
// 			} else {
// 				return false
// 			}
// 		}
// 	}

// 	return st.Len() == 0
// }
