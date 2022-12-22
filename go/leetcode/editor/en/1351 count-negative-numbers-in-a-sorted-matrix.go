package main

//Given a m x n matrix grid which is sorted in non-increasing order both row-
//wise and column-wise, return the number of negative numbers in grid.
//
//
// Example 1:
//
//
//Input: grid = [[4,3,2,-1],[3,2,1,-1],[1,1,-1,-2],[-1,-1,-2,-3]]
//Output: 8
//Explanation: There are 8 negatives number in the matrix.
//
//
// Example 2:
//
//
//Input: grid = [[3,2],[1,0]]
//Output: 0
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 100
// -100 <= grid[i][j] <= 100
//
//
//
//Follow up: Could you find an
//O(n + m) solution?
//
// Related Topics Array Binary Search Matrix 👍 2912 👎 87

//leetcode submit region begin(Prohibit modification and deletion)
// 注意每個 row 的第一個元素已經由小排到大，所以 column 只需要遍歷一次
func countNegatives(grid [][]int) int {
	count := 0
	r, c := len(grid)-1, 0
	for r >= 0 && c < len(grid[r]) {
		if grid[r][c] >= 0 {
			c++
			continue
		}
		count += len(grid[r]) - c
		r--
	}

	return count
}

//leetcode submit region end(Prohibit modification and deletion)
