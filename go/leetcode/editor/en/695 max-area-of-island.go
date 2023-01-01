package main

//You are given an m x n binary matrix grid. An island is a group of 1's (
//representing land) connected 4-directionally (horizontal or vertical.) You may assume
//all four edges of the grid are surrounded by water.
//
// The area of an island is the number of cells with a value 1 in the island.
//
// Return the maximum area of an island in grid. If there is no island, return 0
//.
//
//
// Example 1:
//
//
//Input: grid = [[0,0,1,0,0,0,0,1,0,0,0,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,1,1,
//0,1,0,0,0,0,0,0,0,0],[0,1,0,0,1,1,0,0,1,0,1,0,0],[0,1,0,0,1,1,0,0,1,1,1,0,0],[0,
//0,0,0,0,0,0,0,0,0,1,0,0],[0,0,0,0,0,0,0,1,1,1,0,0,0],[0,0,0,0,0,0,0,1,1,0,0,0,0]
//]
//Output: 6
//Explanation: The answer is not 11, because the island must be connected 4-
//directionally.
//
//
// Example 2:
//
//
//Input: grid = [[0,0,0,0,0,0,0,0]]
//Output: 0
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 50
// grid[i][j] is either 0 or 1.
//
//
// Related Topics Array Depth-First Search Breadth-First Search Union Find
//Matrix 👍 8419 👎 185

//leetcode submit region begin(Prohibit modification and deletion)
func maxAreaOfIsland(grid [][]int) int {
	maxNum := 0
	dirs := [4][2]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}
	rowLen, colLen := len(grid), len(grid[0])
	var maxArea func(row, col int, count *int)
	maxArea = func(row, col int, count *int) {
		if row < 0 || col < 0 || row >= rowLen || col >= colLen || grid[row][col] <= 0 {
			return
		}
		*count++
		maxNum = max(maxNum, *count)
		grid[row][col] *= -1
		for _, dir := range dirs {
			maxArea(row+dir[0], col+dir[1], count)
		}
	}

	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] != 1 {
				continue
			}
			count := 0
			maxArea(row, col, &count)
		}
	}

	return maxNum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
