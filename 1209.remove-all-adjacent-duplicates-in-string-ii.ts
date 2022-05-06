/*
 * @lc app=leetcode id=1209 lang=typescript
 *
 * [1209] Remove All Adjacent Duplicates in String II
 *
 * https://leetcode.com/problems/remove-all-adjacent-duplicates-in-string-ii/description/
 *
 * algorithms
 * Medium (56.11%)
 * Likes:    2978
 * Dislikes: 57
 * Total Accepted:    171.2K
 * Total Submissions: 305.1K
 * Testcase Example:  '"abcd"\n2'
 *
 * You are given a string s and an integer k, a k duplicate removal consists of
 * choosing k adjacent and equal letters from s and removing them, causing the
 * left and the right side of the deleted substring to concatenate together.
 *
 * We repeatedly make k duplicate removals on s until we no longer can.
 *
 * Return the final string after all such duplicate removals have been made. It
 * is guaranteed that the answer is unique.
 *
 *
 * Example 1:
 *
 *
 * Input: s = "abcd", k = 2
 * Output: "abcd"
 * Explanation: There's nothing to delete.
 *
 * Example 2:
 *
 *
 * Input: s = "deeedbbcccbdaa", k = 3
 * Output: "aa"
 * Explanation:
 * First delete "eee" and "ccc", get "ddbbbdaa"
 * Then delete "bbb", get "dddaa"
 * Finally delete "ddd", get "aa"
 *
 * Example 3:
 *
 *
 * Input: s = "pbbcggttciiippooaais", k = 2
 * Output: "ps"
 *
 *
 *
 * Constraints:
 *
 *
 * 1 <= s.length <= 10^5
 * 2 <= k <= 10^4
 * s only contains lower case English letters.
 *
 *
 */
console.log(removeDuplicates("pbbcggttciiippooaais", 2));

// @lc code=start
type DupNode = {
  char: string;
  length: number;
};
function removeDuplicates(s: string, k: number): string {
  const res = Array<string>();
  const dupList = Array<DupNode>();
  dupList.push({ char: "", length: 1 });
  for (const char of s) {
    const lastDupNode: DupNode = dupList[dupList.length - 1];
    if (char == lastDupNode.char) {
      res.push(char);
      lastDupNode.length++;
      if (lastDupNode.length >= k) {
        for (let i = 0; i < k; i++) {
          res.pop();
        }
        lastDupNode.length -= k;
        if (lastDupNode.length <= 0) {
          dupList.pop();
        }
      }
    } else {
      res.push(char);
      dupList.push({ char, length: 1 });
    }
  }

  // for (let i = 1; i < dupList.length; i++) {
  //   for (let j = 0; j < dupList[i].length; j++) {
  //     res.push(dupList[i].char);
  //   }
  // }

  return res.join("");
}
// @lc code=end
