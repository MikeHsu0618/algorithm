package main

import "strings"

//A valid IP address consists of exactly four integers separated by single dots.
// Each integer is between 0 and 255 (inclusive) and cannot have leading zeros.
//
//
// For example, "0.1.2.201" and "192.168.1.1" are valid IP addresses, but "0.011
//.255.245", "192.168.1.312" and "192.168@1.1" are invalid IP addresses.
//
//
// Given a string s containing only digits, return all possible valid IP
//addresses that can be formed by inserting dots into s. You are not allowed to reorder
//or remove any digits in s. You may return the valid IP addresses in any order.
//
//
// Example 1:
//
//
//Input: s = "25525511135"
//Output: ["255.255.11.135","255.255.111.35"]
//
//
// Example 2:
//
//
//Input: s = "0000"
//Output: ["0.0.0.0"]
//
//
// Example 3:
//
//
//Input: s = "101023"
//Output: ["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]
//
//
//
// Constraints:
//
//
// 1 <= s.length <= 20
// s consists of digits only.
//
//
// Related Topics String Backtracking 👍 3306 👎 666

//leetcode submit region begin(Prohibit modification and deletion)
func restoreIpAddresses(s string) []string {
	res := make([]string, 0)
	path := make([]string, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == 4 {
			if start == len(s) {
				res = append(res, strings.Join(path, "."))
			}
			return
		}

		for i := start; i < len(s); i++ {
			if !isValid(s, start, i) {
				continue
			}
			str := s[start : i+1]
			path = append(path, str)
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)
	return res
}

func isValid(s string, start, end int) bool {
	if start > end {
		return false
	}
	// 不允許開頭為零的數字，如 01 但可以允許 0
	if s[start] == '0' && start != end {
		return false
	}

	num := 0
	for i := start; i <= end; i++ {
		if s[i] > '9' || s[i] < '0' {
			return false
		}
		num = num*10 + int(s[i]-'0')
		if num > 255 {
			return false
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
