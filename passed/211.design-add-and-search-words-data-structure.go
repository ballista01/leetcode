/*
 * @lc app=leetcode id=211 lang=golang
 *
 * [211] Design Add and Search Words Data Structure
 */
package main

// @lc code=start
type WordDictionary struct {
	c        byte
	children [26]*WordDictionary
	terminal bool
}

func Constructor() WordDictionary {
	return WordDictionary{byte('@'), [26]*WordDictionary{}, false}
}

func (this *WordDictionary) AddWord(word string) {
	parent := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		cIndex := c - 'a'
		// fmt.Printf("c is %c, c - a = %d - %d = %d\n", c, c, 'a', cIndex)
		if parent.children[cIndex] == nil {
			parent.children[cIndex] = &WordDictionary{c, [26]*WordDictionary{}, false}
		}
		parent = parent.children[cIndex]
	}
	parent.terminal = true
}

func (this *WordDictionary) Search(word string) bool {
	parent := this
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c != '.' {
			cIndex := c - 'a'
			// fmt.Printf("c is %c, c - a = %d - %d = %d\n", c, c, 'a', cIndex)
			if parent.children[cIndex] == nil {
				return false
			}
			parent = parent.children[cIndex]
		} else {
			for j := 0; j < len(parent.children); j++ {
				if parent.children[j] != nil && parent.children[j].Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
	}
	return parent.terminal
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

// Accepted
// 29/29 cases passed (630 ms)
// Your runtime beats 98.69 % of golang submissions
// Your memory usage beats 51.53 % of golang submissions (62.1 MB)
