package main

//Given a string s which consists of lowercase or uppercase letters, return the
//length of the longest palindrome that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a palindrome
//here.
//
//
// Example 1:
//
//
//Input: s = "abccccdd"
//Output: 7
//Explanation: One longest palindrome that can be built is "dccaccd", whose
//length is 7.
//
//
// Example 2:
//
//
//Input: s = "a"
//Output: 1
//Explanation: The longest palindrome that can be built is "a", whose length is
//1.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 2000
// s consists of lowercase and/or uppercase English letters only.
//
//
// Related Topics Hash Table String Greedy 👍 3770 👎 223

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	// 統計每個字母出現的次數
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 關鍵：計算奇數個字母出現的次數，因為奇數個字母一定會有一個落單
	count := 0
	for _, v := range m {
		if v%2 == 1 {
			count++
		}
	}
	// 迴文可以允許一個落單的字母，所以要加回去
	if count > 0 {
		return len(s) - count + 1
	}
	return len(s)
}

//leetcode submit region end(Prohibit modification and deletion)
