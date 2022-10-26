package main

//Given a string s, return the longest palindromic substring in s.
//
// A string is called a palindrome string if the reverse of that string is the
//same as the original string.
//
//
// Example 1:
//
//
//Input: s = "babad"
//Output: "bab"
//Explanation: "aba" is also a valid answer.
//
//
// Example 2:
//
//
//Input: s = "cbbd"
//Output: "bb"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 1000
// s consist of only digits and English letters.
//
//
// Related Topics String Dynamic Programming 👍 22217 👎 1278

//leetcode submit region begin(Prohibit modification and deletion)
// TODO 回來用馬拉車解法再實現一次
func longestPalindrome1(s string) string {
	// 判斷邊界
	if len(s) < 2 {
		return s
	}
	// 宣告起點、最大長度
	var start, maxLen = 0, 0
	// 迴文的特點是利用找到中心點（一個或多個相同字母）
	// 然後以 begin - 1 和 end + 1 向外擴散比對
	// 最終取得該次最長迴文長度，並比較是否刷新紀錄
	// 一開始我們以 i 為中心段去找到擴散
	for i := 0; i < len(s); i++ {
		// 因為 i 為中間字段所以 maxLen/2 大於 len(s) - i 就不用繼續找了
		if len(s)-i <= maxLen/2 {
			break
		}
		// 設立中心點的起點與終點
		begin, end := i, i
		// 從中心點先找到最長重複字符
		for j := end; j < len(s)-1; j++ {
			if s[j] != s[j+1] {
				break
			}
			end++
		}
		// 為下次中間字段做準備
		i = end

		// 開始由中心點兩側開始查找（需注意邊界範圍）
		for f := end; f < len(s)-1 && begin > 0; f++ {
			if s[end+1] != s[begin-1] {
				break
			}
			end++
			begin--
		}
		// 刷新最大值
		if end-begin+1 > maxLen {
			start = begin
			maxLen = end - begin + 1
		}
	}
	return s[start : start+maxLen]
}

//leetcode submit region end(Prohibit modification and deletion)
