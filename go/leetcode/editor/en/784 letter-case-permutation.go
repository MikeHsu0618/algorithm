package main

//Given a string s, you can transform every letter individually to be lowercase
//or uppercase to create another string.
//
// Return a list of all possible strings we could create. Return the output in
//any order.
//
//
// Example 1:
//
//
//Input: s = "a1b2"
//Output: ["a1b2","a1B2","A1b2","A1B2"]
//
//
// Example 2:
//
//
//Input: s = "3z4"
//Output: ["3z4","3Z4"]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 12
// s consists of lowercase English letters, uppercase English letters, and
//digits.
//
//
// Related Topics String Backtracking Bit Manipulation 👍 4013 👎 150

//leetcode submit region begin(Prohibit modification and deletion)
// 大小寫字符轉換寫法
// ch := 'a'
// ch = ch - 32
// fmt.Println(ch)  // 輸出：A
func letterCasePermutation(s string) []string {
	res := make([]string, 0)
	str := make([]byte, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(str) == len(s) {
			res = append(res, string(str))
			return
		}
		if start >= len(s) {
			return
		}
		// 每回合回溯兩次，第一次用原本數值，第二次如果是字母就轉換大小寫

		// 第一次
		str = append(str, s[start])
		backtracking(start + 1)

		// 第二次
		// lowercase -> uppercase
		if s[start] >= 'a' && s[start] <= 'z' {
			str[len(str)-1] -= 32
			backtracking(start + 1)
		}
		// uppercase -> lowercase
		if s[start] >= 'A' && s[start] <= 'Z' {
			str[len(str)-1] += 32
			backtracking(start + 1)
		}
		str = str[:len(str)-1]
	}
	backtracking(0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
