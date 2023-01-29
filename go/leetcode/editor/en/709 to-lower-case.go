package main

//Given a string s, return the string after replacing every uppercase letter
//with the same lowercase letter.
//
//
// Example 1:
//
//
//Input: s = "Hello"
//Output: "hello"
//
//
// Example 2:
//
//
//Input: s = "here"
//Output: "here"
//
//
// Example 3:
//
//
//Input: s = "LOVELY"
//Output: "lovely"
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 100
// s consists of printable ASCII characters.
//
//
// Related Topics String 👍 1433 👎 2560

//leetcode submit region begin(Prohibit modification and deletion)
func toLowerCase(str string) string {
	s := make([]byte, len(str))

	for i := 0; i < len(str); i++ {
		if str[i] >= 'A' && str[i] <= 'Z' {
			s[i] = str[i] + byte(32)
			continue
		}
		s[i] = str[i]
	}

	return string(s)
}

//leetcode submit region end(Prohibit modification and deletion)
