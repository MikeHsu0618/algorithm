package main

//Given a string s, reverse the order of characters in each word within a
//sentence while still preserving whitespace and initial word order.
//
//
// Example 1:
// Input: s = "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
//
// Example 2:
// Input: s = "God Ding"
//Output: "doG gniD"
//
//
// Constraints:
//
//
// 1 <= s.length <= 5 * 10⁴
// s contains printable ASCII characters.
// s does not contain any leading or trailing spaces.
// There is at least one word in s.
// All the words in s are separated by a single space.
//
//
// Related Topics Two Pointers String 👍 4434 👎 216

//leetcode submit region begin(Prohibit modification and deletion)
func reverseWords(s string) string {
	slow, fast := 0, 0
	ss := []byte(s)
	for fast < len(ss) {
		if s[fast] == ' ' {
			reverse(ss[slow:fast])
			slow = fast + 1
		}

		if fast+1 == len(ss) {
			reverse(ss[slow : fast+1])
		}

		fast++
	}
	return string(ss)
}

func reverse(s []byte) {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
