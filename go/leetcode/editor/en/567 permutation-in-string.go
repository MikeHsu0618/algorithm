package main

//Given two strings s1 and s2, return true if s2 contains a permutation of s1,
//or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of
//s2.
//
//
// Example 1:
//
//
//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
//
//
// Example 2:
//
//
//Input: s1 = "ab", s2 = "eidboaoo"
//Output: false
//
//
//
// Constraints:
//
//
// 1 <= s1.length, s2.length <= 10⁴
// s1 and s2 consist of lowercase English letters.
//
//
// Related Topics Hash Table Two Pointers String Sliding Window 👍 7820 👎 258

//leetcode submit region begin(Prohibit modification and deletion)
// 解法： 用 slice window 找出連續的在 s2 中的 s1(全排序)子字串
func checkInclusion(s1 string, s2 string) bool {
	start, count := 0, [26]byte{}
	// 將 s1 放進 map 或特製的 array 中
	for i := range s1 {
		count[s1[i]-'a']++
	}

	// 使用 slice window 像 count 中的字母表相減，直到表歸零及返回 true
	for end := range s2 {
		count[s2[end]-'a']--
		if count == [26]byte{} {
			return true
		}

		// 當長度開始大於 s2 時，開始移動 start 指針並且復原字母表
		if end+1 >= len(s1) {
			count[s2[start]-'a']++
			start++
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
