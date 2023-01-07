package main

//Given an integer n, return true if it is a power of two. Otherwise, return
//false.
//
// An integer n is a power of two, if there exists an integer x such that n == 2
//ˣ.
//
//
// Example 1:
//
//
//Input: n = 1
//Output: true
//Explanation: 2⁰ = 1
//
//
// Example 2:
//
//
//Input: n = 16
//Output: true
//Explanation: 2⁴ = 16
//
//
// Example 3:
//
//
//Input: n = 3
//Output: false
//
//
//
// Constraints:
//
//
// -2³¹ <= n <= 2³¹ - 1
//
//
//
//Follow up: Could you solve it without loops/recursion?
//
// Related Topics Math Bit Manipulation Recursion 👍 4626 👎 341

//leetcode submit region begin(Prohibit modification and deletion)
func isPowerOfTwo(n int) bool {
	// ＆ 為二進位下的 AND 位元運算子
	// 100 & 011 = 000 -> 4 & 3 = 0
	// 就是將兩個數字的二進位表示的每一位元分別相and，最終的結果就是 000。
	return n > 0 && n&(n-1) == 0
}

//leetcode submit region end(Prohibit modification and deletion)
