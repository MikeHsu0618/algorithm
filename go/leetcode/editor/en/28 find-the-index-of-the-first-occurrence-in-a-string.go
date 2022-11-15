package main

//Given two strings needle and haystack, return the index of the first
//occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//
// Example 1:
//
//
//Input: haystack = "sadbutsad", needle = "sad"
//Output: 0
//Explanation: "sad" occurs at index 0 and 6.
//The first occurrence is at index 0, so we return 0.
//
//
// Example 2:
//
//
//Input: haystack = "leetcode", needle = "leeto"
//Output: -1
//Explanation: "leeto" did not occur in "leetcode", so we return -1.
//
//
//
// Constraints:
//
//
// 1 <= haystack.length, needle.length <= 10⁴
// haystack and needle consist of only lowercase English characters.
//
//
// Related Topics Two Pointers String String Matching 👍 845 👎 66

//leetcode submit region begin(Prohibit modification and deletion)
// 解法三： KMP (TODO)
func strStr(haystack string, needle string) int {

}

// 解法二(優化暴力解)：之接遍歷 haystack[i:len(needle)] == need
func strStr2(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i:len(needle)] == needle {
			return i
		}
	}
	return -1
}

// 解法一(暴力解)：遍歷字串查找，只需要找到第 len(haystack) - len(needle) + 1 個即可
func strStr1(haystack string, needle string) int {
	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] != needle[0] {
			continue
		}
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
