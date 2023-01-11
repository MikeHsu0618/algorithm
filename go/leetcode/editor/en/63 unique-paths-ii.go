package main

//You are given an m x n integer array grid. There is a robot initially located
//at the top-left corner (i.e., grid[0][0]). The robot tries to move to the
//bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down
//or right at any point in time.
//
// An obstacle and space are marked as 1 or 0 respectively in grid. A path that
//the robot takes cannot include any square that is an obstacle.
//
// Return the number of possible unique paths that the robot can take to reach
//the bottom-right corner.
//
// The testcases are generated so that the answer will be less than or equal to
//2 * 10⁹.
//
//
// Example 1:
//
//
//Input: obstacleGrid = [[0,0,0],[0,1,0],[0,0,0]]
//Output: 2
//Explanation: There is one obstacle in the middle of the 3x3 grid above.
//There are two ways to reach the bottom-right corner:
//1. Right -> Right -> Down -> Down
//2. Down -> Down -> Right -> Right
//
//
// Example 2:
//
//
//Input: obstacleGrid = [[0,1],[0,0]]
//Output: 1
//
//
//
// Constraints:
//
//
// m == obstacleGrid.length
// n == obstacleGrid[i].length
// 1 <= m, n <= 100
// obstacleGrid[i][j] is 0 or 1.
//
//
// Related Topics Array Dynamic Programming Matrix 👍 6439 👎 423

//leetcode submit region begin(Prohibit modification and deletion)
// 加入障礙物，代表遇到障礙物後只能往下走且後面的路徑都不會經過，因為機器人只能往右或往下
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 如果 obstacleGrid[i][j] 這個點是障礙物
			if obstacleGrid[i][j] == 0 {
				// 否則我們需要計算當前點可以到達的路徑數
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
