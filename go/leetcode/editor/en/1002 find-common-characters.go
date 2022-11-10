package main

//Given a string array words, return an array of all characters that show up in
//all strings within the words (including duplicates). You may return the answer
//in any order.
//
//
// Example 1:
// Input: words = ["bella","label","roller"]
//Output: ["e","l","l"]
//
// Example 2:
// Input: words = ["cool","lock","cook"]
//Output: ["c","o"]
//
//
// Constraints:
//
//
// 1 <= words.length <= 100
// 1 <= words[i].length <= 100
// words[i] consists of lowercase English letters.
//
//
// Related Topics Array Hash Table String 👍 2700 👎 222

//leetcode submit region begin(Prohibit modification and deletion)
// 找到所有字串中共通的字母
// 思路：以第一個字串依據，並遍歷所有後續字串與第一個比較，比較共通字母數量（取小者），或刪除不共通字母
func commonChars(words []string) []string {
	res := make([]string, 0)
	m := make(map[byte]int)
	for i := 0; i < len(words[0]); i++ {
		m[words[0][i]]++
	}
	for i := 1; i < len(words); i++ {
		m2 := make(map[byte]int)
		for j := 0; j < len(words[i]); j++ {
			m2[words[i][j]]++
		}
		for i, _ := range m {
			if _, ok := m2[i]; !ok {
				delete(m, i)
				continue
			}
			if m[i] > m2[i] {
				m[i] = m2[i]
			}
		}
	}
	for i, v := range m {
		for j := 0; j < v; j++ {
			res = append(res, string(i))
		}
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
