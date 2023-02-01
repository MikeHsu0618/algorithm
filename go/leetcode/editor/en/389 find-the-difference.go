package main

//You are given two strings s and t.
//
// String t is generated by random shuffling string s and then add one more
//letter at a random position.
//
// Return the letter that was added to t.
//
//
// Example 1:
//
//
//Input: s = "abcd", t = "abcde"
//Output: "e"
//Explanation: 'e' is the letter that was added.
//
//
// Example 2:
//
//
//Input: s = "", t = "y"
//Output: "y"
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 1000
// t.length == s.length + 1
// s and t consist of lowercase English letters.
//
//
// Related Topics Hash Table String Bit Manipulation Sorting 👍 3564 👎 409

//leetcode submit region begin(Prohibit modification and deletion)
func findTheDifference(s string, t string) byte {
	m := make(map[byte]int, 0)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}

	for i := 0; i < len(t); i++ {
		if v, ok := m[t[i]]; ok && v > 0 {
			m[t[i]]--
			continue
		}

		return t[i]
	}

	return t[len(t)-1]
}

//leetcode submit region end(Prohibit modification and deletion)