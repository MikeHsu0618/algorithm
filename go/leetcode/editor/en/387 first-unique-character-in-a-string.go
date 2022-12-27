package main

//Given a string s, find the first non-repeating character in it and return its
//index. If it does not exist, return -1.
//
//
// Example 1:
// Input: s = "leetcode"
//Output: 0
//
// Example 2:
// Input: s = "loveleetcode"
//Output: 2
//
// Example 3:
// Input: s = "aabb"
//Output: -1
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁵
// s consists of only lowercase English letters.
//
//
// Related Topics Hash Table String Queue Counting 👍 7122 👎 239

//leetcode submit region begin(Prohibit modification and deletion)
// 使用 ASCII 處理字母的解法
func firstUniqChar(s string) int {
	m := make([]int, 26)
	for i := 0; i < len(s); i++ {
		m[s[i]-'a']++
	}
	for i := 0; i < len(s); i++ {
		if m[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}

// 優化解法一
func firstUniqChar2(s string) int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok && v == 1 {
			return i
		}
	}

	return -1
}

func firstUniqChar1(s string) int {
	m := make(map[byte]int)
	q := make([]int, len(s))
	// 將字串同步塞進 hash map 和 queue
	for i := 0; i < len(s); i++ {
		m[s[i]]++
		q[i] = i
	}
	// 將 queue 依序取出來查看，找到第一個不重複的
	for i := 0; i < len(q); i++ {
		if v, ok := m[s[q[i]]]; ok && v == 1 {
			return q[i]
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
