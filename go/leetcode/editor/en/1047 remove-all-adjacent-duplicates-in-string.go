package main

//You are given a string s consisting of lowercase English letters. A duplicate
//removal consists of choosing two adjacent and equal letters and removing them.
//
// We repeatedly make duplicate removals on s until we no longer can.
//
// Return the final string after all such duplicate removals have been made. It
//can be proven that the answer is unique.
//
//
// Example 1:
//
//
//Input: s = "abbaca"
//Output: "ca"
//Explanation:
//For example, in "abbaca" we could remove "bb" since the letters are adjacent
//and equal, and this is the only possible move.  The result of this move is that
//the string is "aaca", of which only "aa" is possible, so the final string is
//"ca".
//
//
// Example 2:
//
//
//Input: s = "azxxzy"
//Output: "ay"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁵
// s consists of lowercase English letters.
//
//
// Related Topics String Stack 👍 5145 👎 203

//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if i == 0 || len(stack) < 2 || stack[len(stack)-1] != stack[len(stack)-2] {
			continue
		}
		stack = stack[:len(stack)-2]
	}
	return string(stack)
}

//leetcode submit region end(Prohibit modification and deletion)
