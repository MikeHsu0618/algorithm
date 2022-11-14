package main

//Given a string s and an integer k, reverse the first k characters for every 2
//k characters counting from the start of the string.
//
// If there are fewer than k characters left, reverse all of them. If there are
//less than 2k but greater than or equal to k characters, then reverse the first
//k characters and leave the other as original.
//
//
// Example 1:
// Input: s = "abcdefg", k = 2
//Output: "bacdfeg"
//
// Example 2:
// Input: s = "abcd", k = 2
//Output: "bacd"
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s consists of only lowercase English letters.
// 1 <= k <= 10⁴
//
//
// Related Topics Two Pointers String 👍 1339 👎 2859

//leetcode submit region begin(Prohibit modification and deletion)
func reverseStr(s string, k int) string {
	// 需要轉換成 []byte 才能控制 string（string 是不變的）
	ss := []byte(s)
	// 每次往前推進 2*k 個區間
	for i := 0; i < len(s); i += 2 * k {
		// 查看後面要被 reverse 的空間
		if i+k <= len(s) {
			reverse1(ss[i : i+k])
			continue
		}
		// 如果沒空間就將剩下的全部調換
		reverse1(ss[i:len(s)])
	}
	return string(ss)
}

func reverse1(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
