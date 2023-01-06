package main

//Given a triangle array, return the minimum path sum from top to bottom.
//
// For each step, you may move to an adjacent number of the row below. More
//formally, if you are on index i on the current row, you may move to either index i
//or index i + 1 on the next row.
//
//
// Example 1:
//
//
//Input: triangle = [[2],[3,4],[6,5,7],[4,1,8,3]]
//Output: 11
//Explanation: The triangle looks like:
//   2
//  3 4
// 6 5 7
//4 1 8 3
//The minimum path sum from top to bottom is 2 + 3 + 5 + 1 = 11 (underlined
//above).
//
//
// Example 2:
//
//
//Input: triangle = [[-10]]
//Output: -10
//
//
//
// Constraints:
//
//
// 1 <= triangle.length <= 200
// triangle[0].length == 1
// triangle[i].length == triangle[i - 1].length + 1
// -10⁴ <= triangle[i][j] <= 10⁴
//
//
//
//Follow up: Could you do this using only
//O(n) extra space, where
//n is the total number of rows in the triangle?
//
// Related Topics Array Dynamic Programming 👍 7418 👎 455

//leetcode submit region begin(Prohibit modification and deletion)
// 從下面往上增加，最後會到 triangle[0][0] 即為最小路徑總和
func minimumTotal(triangle [][]int) int {
	for row := len(triangle) - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			triangle[row][col] += min(triangle[row+1][col], triangle[row+1][col+1])
		}
	}

	return triangle[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
