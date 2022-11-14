package main

//Given an input string s, reverse the order of the words.
//
// A word is defined as a sequence of non-space characters. The words in s will
//be separated by at least one space.
//
// Return a string of the words in reverse order concatenated by a single space.
//
//
// Note that s may contain leading or trailing spaces or multiple spaces
//between two words. The returned string should only have a single space separating the
//words. Do not include any extra spaces.
//
//
// Example 1:
//
//
//Input: s = "the sky is blue"
//Output: "blue is sky the"
//
//
// Example 2:
//
//
//Input: s = "  hello world  "
//Output: "world hello"
//Explanation: Your reversed string should not contain leading or trailing
//spaces.
//
//
// Example 3:
//
//
//Input: s = "a good   example"
//Output: "example good a"
//Explanation: You need to reduce multiple spaces between two words to a single
//space in the reversed string.
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 10⁴
// s contains English letters (upper-case and lower-case), digits, and spaces '
//'.
// There is at least one word in s.
//
//
//
// Follow-up: If the string data type is mutable in your language, can you
//solve it in-place with O(1) extra space?
//
// Related Topics Two Pointers String 👍 5015 👎 4408

//leetcode submit region begin(Prohibit modification and deletion)
// 解法：整體思路為 移除多餘空白->反轉字串->反轉字詞
// 移除多餘空格 : "the sky is blue"
// 字符串反轉："eulb si yks eht"
// 單字反轉："blue is sky the"
func reverseWords(s string) string {
	ss := []byte(s)
	ss = removeExtraSpaces(ss)
	reverse2(ss, 0, len(ss)-1)
	slow := 0
	for i := 0; i < len(ss); i++ {
		if ss[i] != ' ' && i != len(ss)-1 {
			continue
		}
		fast := i - 1
		if i == len(ss)-1 {
			fast = i
		}
		reverse2(ss, slow, fast)
		slow = i + 1
	}
	return string(ss)
}

// 在陣列中移除元素，通常會使用快慢指針
func removeExtraSpaces(s []byte) []byte {
	slow := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		if slow != 0 {
			s[slow] = ' '
			slow++
		}
		for i < len(s) && s[i] != ' ' {
			s[slow] = s[i]
			slow++
			i++
		}
	}
	return s[:slow]
}

func reverse2(s []byte, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
