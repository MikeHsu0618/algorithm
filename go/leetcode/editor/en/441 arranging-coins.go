package main

//You have n coins and you want to build a staircase with these coins. The
//staircase consists of k rows where the iᵗʰ row has exactly i coins. The last row of
//the staircase may be incomplete.
//
// Given the integer n, return the number of complete rows of the staircase you
//will build.
//
//
// Example 1:
//
//
//Input: n = 5
//Output: 2
//Explanation: Because the 3ʳᵈ row is incomplete, we return 2.
//
//
// Example 2:
//
//
//Input: n = 8
//Output: 3
//Explanation: Because the 4ᵗʰ row is incomplete, we return 3.
//
//
//
// Constraints:
//
//
// 1 <= n <= 2³¹ - 1
//
//
// Related Topics Math Binary Search 👍 2896 👎 1154

//leetcode submit region begin(Prohibit modification and deletion)
func arrangeCoins(n int) int {
	l, r := 1, n
	for l <= r {
		mid := (l + r) / 2
		if mid*(mid+1)/2 <= n {
			l = mid + 1
			continue
		}
		r = mid - 1
	}
	return r
}

//leetcode submit region end(Prohibit modification and deletion)
