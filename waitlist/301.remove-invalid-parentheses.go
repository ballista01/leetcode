/*
 * @lc app=leetcode id=301 lang=golang
 *
 * [301] Remove Invalid Parentheses
 */
package main
import "container/list"
// @lc code=start
func removeInvalidParentheses(s string) []string {
	st := list.New()
	bal := 0
	res := make([]string, 0)

	if s[0] == ')' {
		bal --
		s=s[1:]
	}
	if s[len(s)-1] == '(' {
		bal ++
		s=s[:len(s)-1]
	}

	for i:=0;i<len(s);i++{
		if s[i] == '(' {
			bal --
		}
		if s[i] == ')' {
			bal ++
		}
	}

	last := byte('@')
	for i:=0; i<len(s); i++{
		if last != s[i]{
			if bal<0 {
				res = append(res, s[])
			}
		}
	}
}

// @lc code=end
