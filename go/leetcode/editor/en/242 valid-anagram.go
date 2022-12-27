package main

//Given two strings s and t, return true if t is an anagram of s, and false
//otherwise.
//
// An Anagram is a word or phrase formed by rearranging the letters of a
//different word or phrase, typically using all the original letters exactly once.
//
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
//Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
//Output: false
//
//
// Constraints:
//
//
// 1 <= s.length, t.length <= 5 * 10⁴
// s and t consist of lowercase English letters.
//
//
//
// Follow up: What if the inputs contain Unicode characters? How would you
//adapt your solution to such a case?
//
// Related Topics Hash Table String Sorting 👍 7445 👎 247

//leetcode submit region begin(Prohibit modification and deletion)
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 遍歷 s 加進 hash map 中
	m := make(map[int]int, 0)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'a')]++
	}
	// 遍歷 t 來比較 hash map，當數值歸零時需刪除該元素
	for i := 0; i < len(t); i++ {
		if _, ok := m[int(t[i]-'a')]; ok {
			m[int(t[i]-'a')]--
			if m[int(t[i]-'a')] == 0 {
				delete(m, int(t[i]-'a'))
			}
			continue
		}

		return false
	}
	// 比較 hash map 是否剛好歸零
	return len(m) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
