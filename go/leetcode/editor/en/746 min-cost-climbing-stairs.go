package main

//You are given an integer array cost where cost[i] is the cost of iᵗʰ step on
//a staircase. Once you pay the cost, you can either climb one or two steps.
//
// You can either start from the step with index 0, or the step with index 1.
//
// Return the minimum cost to reach the top of the floor.
//
//
// Example 1:
//
//
//Input: cost = [10,15,20]
//Output: 15
//Explanation: You will start at index 1.
//- Pay 15 and climb two steps to reach the top.
//The total cost is 15.
//
//
// Example 2:
//
//
//Input: cost = [1,100,1,1,1,100,1,1,100,1]
//Output: 6
//Explanation: You will start at index 0.
//- Pay 1 and climb two steps to reach index 2.
//- Pay 1 and climb two steps to reach index 4.
//- Pay 1 and climb two steps to reach index 6.
//- Pay 1 and climb one step to reach index 7.
//- Pay 1 and climb two steps to reach index 9.
//- Pay 1 and climb one step to reach the top.
//The total cost is 6.
//
//
//
// Constraints:
//
//
// 2 <= cost.length <= 1000
// 0 <= cost[i] <= 999
//
//
// Related Topics Array Dynamic Programming 👍 8610 👎 1345

//leetcode submit region begin(Prohibit modification and deletion)
// 優化
func minCostClimbingStairs(cost []int) int {
	n := len(cost)

	dp := make([]int, n)
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < n; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])
	}
	return min(dp[n-1], dp[n-2])
}

// DP
func minCostClimbingStairs1(cost []int) int {
	// dp[i]的定義：到達第i台階所花費的最少體力為 dp[i]
	dp := make(map[int]int, 0)
	// 默認第一步都不用花體力
	dp[0], dp[1] = 0, 0
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	return dp[len(cost)]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
