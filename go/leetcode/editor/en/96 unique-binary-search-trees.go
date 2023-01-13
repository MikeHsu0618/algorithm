package main

//Given an integer n, return the number of structurally unique BST's (binary
//search trees) which has exactly n nodes of unique values from 1 to n.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: 5
//
//
// Example 2:
//
//
//Input: n = 1
//Output: 1
//
//
//
// Constraints:
//
//
// 1 <= n <= 19
//
//
// Related Topics Math Dynamic Programming Tree Binary Search Tree Binary Tree ?
//? 8639 👎 345

//leetcode submit region begin(Prohibit modification and deletion)
/**
  dp[i]: i个节点对应的种树
  dp[0]: -1; 无意义；
  dp[1]: 1;
  ...
  dp[i]: 2 * dp[i - 1] +
      (dp[1] * dp[i - 2] + dp[2] * dp[i - 3] + ... + dp[i - 2] * dp[1]); 从1加到i-2
*/
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}

	return dp[n]
}

//leetcode submit region end(Prohibit modification and deletion)
