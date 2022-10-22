package main

//Given a string s, find the length of the longest substring without repeating
//characters.
//
//
// Example 1:
//
//
//Input: s = "abcabcbb"
//Output: 3
//Explanation: The answer is "abc", with the length of 3.
//
//
// Example 2:
//
//
//Input: s = "bbbbb"
//Output: 1
//Explanation: The answer is "b", with the length of 1.
//
//
// Example 3:
//
//
//Input: s = "pwwkew"
//Output: 3
//Explanation: The answer is "wke", with the length of 3.
//Notice that the answer must be a substring, "pwke" is a subsequence and not a
//substring.
//
//
//
// Constraints:
//
//
// 0 <= s.length <= 5 * 10⁴
// s consists of English letters, digits, symbols and spaces.
//
//
// Related Topics Hash Table String Sliding Window 👍 29651 👎 1259

//leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	m := make(map[byte]int)
	start, maxLen := 0, 0
	for i := 0; i < len(s); i++ {
		// 檢查 hash table 中以有在 start 後重複出現的字母
		if v, ok := m[s[i]]; ok && v >= start {
			// 有的話就將起點設為該處後面
			start = v + 1
		}
		// 將每次字母出現的位置都記錄下來
		m[s[i]] = i
		// 確認這次重複前的長度是否比之前的紀錄的最大長度還長
		if (i - start + 1) > maxLen {
			maxLen = i - start + 1
		}
		// 如果 start 加上 maxLen 已經超過字串長度可以提早返回
		if start+maxLen > len(s) {
			break
		}
	}
	return maxLen
}

//leetcode submit region end(Prohibit modification and deletion)
