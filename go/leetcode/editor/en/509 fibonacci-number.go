package main

//The Fibonacci numbers, commonly denoted F(n) form a sequence, called the
//Fibonacci sequence, such that each number is the sum of the two preceding ones,
//starting from 0 and 1. That is,
//
//
//F(0) = 0, F(1) = 1
//F(n) = F(n - 1) + F(n - 2), for n > 1.
//
//
// Given n, calculate F(n).
//
//
// Example 1:
//
//
//Input: n = 2
//Output: 1
//Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//
//
// Example 2:
//
//
//Input: n = 3
//Output: 2
//Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//
//
// Example 3:
//
//
//Input: n = 4
//Output: 3
//Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
//
//
//
// Constraints:
//
//
// 0 <= n <= 30
//
//
// Related Topics Math Dynamic Programming Recursion Memoization 👍 5881 👎 302

//leetcode submit region begin(Prohibit modification and deletion)
// 動態 -> 並且存下數值 -> 空間換取時間
func fib(n int) int {
	m := make(map[int]int, 0)

	if n <= 1 {
		return n
	}

	m[0] = 0
	m[1] = 1
	for i := 2; i <= n; i++ {
		sum := m[0] + m[1]
		m[0] = m[1]
		m[1] = sum
	}

	return m[1]
}

// 遞迴寫法：時間複雜度為 O(2^n) 較差
func fib1(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

//leetcode submit region end(Prohibit modification and deletion)
