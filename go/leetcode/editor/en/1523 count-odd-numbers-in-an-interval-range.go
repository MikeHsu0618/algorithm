package main

//Given two non-negative integers low and high. Return the count of odd numbers
//between low and high (inclusive).
//
//
// Example 1:
//
//
//Input: low = 3, high = 7
//Output: 3
//Explanation: The odd numbers between 3 and 7 are [3,5,7].
//
// Example 2:
//
//
//Input: low = 8, high = 10
//Output: 1
//Explanation: The odd numbers between 8 and 10 are [9].
//
//
// Constraints:
//
//
// 0 <= low <= high <= 10^9
//
//
// Related Topics Math 👍 1305 👎 90

//leetcode submit region begin(Prohibit modification and deletion)
// 一行解
func countOdds(low int, high int) int {
	return (high+1)/2 - low/2
}

// 暴力解
func countOdds1(low int, high int) int {
	count := 0
	for i := low; i <= high; i++ {
		if i%2 == 1 {
			count++
		}
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
