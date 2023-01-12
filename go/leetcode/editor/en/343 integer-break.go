package main

//Given an integer n, break it into the sum of k positive integers, where k >= 2
//, and maximize the product of those integers.
//
// Return the maximum product you can get.
//
//
// Example 1:
//
//
//Input: n = 2
//Output: 1
//Explanation: 2 = 1 + 1, 1 × 1 = 1.
//
//
// Example 2:
//
//
//Input: n = 10
//Output: 36
//Explanation: 10 = 3 + 3 + 4, 3 × 3 × 4 = 36.
//
//
//
// Constraints:
//
//
// 2 <= n <= 58
//
//
// Related Topics Math Dynamic Programming 👍 3422 👎 355

//leetcode submit region begin(Prohibit modification and deletion)
func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}

	dp := make([]int, n+1)

	dp[2] = 1
	for i := 3; i <= n; i++ {
		// i 可以拆為 i - j 和 j 。
		for j := 1; j <= i/2; j++ {
			// 由于需要最大值，故需要通过j遍历所有存在的值，取其中最大的值作为当前i的最大值，在求最大值的时候，
			// 一个是j与i-j相乘(拆成兩個)，一个是j与dp[i-j]（拆成兩個以上）.
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}

	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
