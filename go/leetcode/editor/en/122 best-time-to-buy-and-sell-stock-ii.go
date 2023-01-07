package main

//You are given an integer array prices where prices[i] is the price of a given
//stock on the iᵗʰ day.
//
// On each day, you may decide to buy and/or sell the stock. You can only hold
//at most one share of the stock at any time. However, you can buy it then
//immediately sell it on the same day.
//
// Find and return the maximum profit you can achieve.
//
//
// Example 1:
//
//
//Input: prices = [7,1,5,3,6,4]
//Output: 7
//Explanation: Buy on day 2 (price = 1) and sell on day 3 (price = 5), profit =
//5-1 = 4.
//Then buy on day 4 (price = 3) and sell on day 5 (price = 6), profit = 6-3 = 3.
//
//Total profit is 4 + 3 = 7.
//
//
// Example 2:
//
//
//Input: prices = [1,2,3,4,5]
//Output: 4
//Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit =
//5-1 = 4.
//Total profit is 4.
//
//
// Example 3:
//
//
//Input: prices = [7,6,4,3,1]
//Output: 0
//Explanation: There is no way to make a positive profit, so we never buy the
//stock to achieve the maximum profit of 0.
//
//
//
// Constraints:
//
//
// 1 <= prices.length <= 3 * 10⁴
// 0 <= prices[i] <= 10⁴
//
//
// Related Topics Array Dynamic Programming Greedy 👍 10079 👎 2491

//leetcode submit region begin(Prohibit modification and deletion)
// 動態規劃：dp[i][0]表示在状态i不持有股票的现金，dp[i][1]为持有股票的现金
func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0], dp[0][1] = 0, -prices[0]
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][0]-prices[i], dp[i-1][1])
	}

	return dp[len(prices)-1][0]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 貪心法：找出每天最佳獲利(局部最佳獲利)
func maxProfit1(prices []int) int {
	maxProfit := 0

	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] > 0 {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}

//leetcode submit region end(Prohibit modification and deletion)
