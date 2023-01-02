package main

//Given a string s, partition s such that every substring of the partition is a
//palindrome. Return all possible palindrome partitioning of s.
//
//
// Example 1:
// Input: s = "aab"
//Output: [["a","a","b"],["aa","b"]]
//
// Example 2:
// Input: s = "a"
//Output: [["a"]]
//
//
// Constraints:
//
//
// 1 <= s.length <= 16
// s contains only lowercase English letters.
//
//
// Related Topics String Dynamic Programming Backtracking 👍 8966 👎 284

//leetcode submit region begin(Prohibit modification and deletion)
func partition(s string) [][]string {
	res := make([][]string, 0)
	path := make([]string, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		if start == len(s) {
			res = append(res, append([]string{}, path...))
			return
		}
		// 以 start 為起點，檢查 start 到 end 的區間是否為迴文
		for i := start; i < len(s); i++ {
			if !isPalindrome(s, start, i) {
				continue
			}
			path = append(path, s[start:i+1])
			// 起點內縮，
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return res
}

// [["a","a","b"],["aa","b"]] 本身或者頭尾依序相同都是迴文
func isPalindrome(s string, start, end int) bool {
	for i, j := start, end; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
