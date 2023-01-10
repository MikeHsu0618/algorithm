package main

//You are given an array prices where prices[i] is the price of a given stock
//on the iᵗʰ day, and an integer fee representing a transaction fee.
//
// Find the maximum profit you can achieve. You may complete as many
//transactions as you like, but you need to pay the transaction fee for each transaction.
//
// Note: You may not engage in multiple transactions simultaneously (i.e., you
//must sell the stock before you buy again).
//
//
// Example 1:
//
//
//Input: prices = [1,3,2,8,4,9], fee = 2
//Output: 8
//Explanation: The maximum profit can be achieved by:
//- Buying at prices[0] = 1
//- Selling at prices[3] = 8
//- Buying at prices[4] = 4
//- Selling at prices[5] = 9
//The total profit is ((8 - 1) - 2) + ((9 - 4) - 2) = 8.
//
//
// Example 2:
//
//
//Input: prices = [1,3,7,5,10,3], fee = 3
//Output: 6
//
//
//
// Constraints:
//
//
// 1 <= prices.length <= 5 * 10⁴
// 1 <= prices[i] < 5 * 10⁴
// 0 <= fee < 5 * 10⁴
//
//
// Related Topics Array Dynamic Programming Greedy 👍 4942 👎 121

//leetcode submit region begin(Prohibit modification and deletion)
func maxProfit(prices []int, fee int) int {
	// dp[i][0]第i天持有股票所剩的最多现金
	// dp[i][1]第i天持有的最多现金
	dp := make([][2]int, len(prices))
	dp[0][0] -= prices[0]
	dp[0][1] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}
	return max(dp[len(prices)-1][0], dp[len(prices)-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxProfit1(prices []int, fee int) int {
	res := 0
	// 第一天買入
	min := prices[0]

	for i := 0; i < len(prices); i++ {
		// 當前價格小於最低價則買入
		if prices[i] < min {
			min = prices[i]
		}

		// 如果以當前價格賣出會虧本，則不賣
		if prices[i] >= min && prices[i] <= min+fee {
			continue
		}
		// 賣出
		if prices[i] > min+fee {
			res += prices[i] - min - fee
			// 更新最小值（如果还在收获利润的区间里，表示并不是真正的卖出，而计算利润每次都要减去手续费
			// 所以要让minBuy = prices[i] - fee这样在明天收获利润的时候，才不会多减一次手续费！）
			min = prices[i] - fee
		}
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
