package main

//You are given an m x n grid where each cell can have one of three values:
//
//
// 0 representing an empty cell,
// 1 representing a fresh orange, or
// 2 representing a rotten orange.
//
//
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten
//orange becomes rotten.
//
// Return the minimum number of minutes that must elapse until no cell has a
//fresh orange. If this is impossible, return -1.
//
//
// Example 1:
//
//
//Input: grid = [[2,1,1],[1,1,0],[0,1,1]]
//Output: 4
//
//
// Example 2:
//
//
//Input: grid = [[2,1,1],[0,1,1],[1,0,1]]
//Output: -1
//Explanation: The orange in the bottom left corner (row 2, column 0) is never
//rotten, because rotting only happens 4-directionally.
//
//
// Example 3:
//
//
//Input: grid = [[0,2]]
//Output: 0
//Explanation: Since there are already no fresh oranges at minute 0, the answer
//is just 0.
//
//
//
// Constraints:
//
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 10
// grid[i][j] is 0, 1, or 2.
//
//
// Related Topics Array Breadth-First Search Matrix 👍 9328 👎 324

//leetcode submit region begin(Prohibit modification and deletion)
// 解法一：BFS
func orangesRotting(grid [][]int) int {
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	Fresh, Rotten := 1, 2
	res := 0
	fresh := 0
	// 將 rotten 放進 queue 裡面，並紀錄 fresh 的數量
	rowLen, colLen := len(grid), len(grid[0])
	queue := make([][2]int, 0)
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if grid[row][col] == Fresh {
				fresh++
			}
			if grid[row][col] == Rotten {
				queue = append(queue, [2]int{row, col})
			}
		}
	}

	if fresh == 0 {
		return res
	}

	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			rotten := queue[i]
			for _, dir := range dirs {
				row, col := rotten[0]+dir[0], rotten[1]+dir[1]
				if row < 0 || row >= rowLen || col < 0 || col >= colLen || grid[row][col] != Fresh {
					continue
				}
				grid[row][col] = Rotten
				queue = append(queue, [2]int{row, col})
				fresh--
			}
		}
		queue = queue[length:]

		res++
		if fresh == 0 {
			return res
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
