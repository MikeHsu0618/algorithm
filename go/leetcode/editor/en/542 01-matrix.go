package main

import "math"

//Given an m x n binary matrix mat, return the distance of the nearest 0 for
//each cell.
//
// The distance between two adjacent cells is 1.
//
//
// Example 1:
//
//
//Input: mat = [[0,0,0],[0,1,0],[0,0,0]]
//Output: [[0,0,0],[0,1,0],[0,0,0]]
//
//
// Example 2:
//
//
//Input: mat = [[0,0,0],[0,1,0],[1,1,1]]
//Output: [[0,0,0],[0,1,0],[1,2,1]]
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 10⁴
// 1 <= m * n <= 10⁴
// mat[i][j] is either 0 or 1.
// There is at least one 0 in mat.
//
//
// Related Topics Array Dynamic Programming Breadth-First Search Matrix 👍 6303
//👎 304

//leetcode submit region begin(Prohibit modification and deletion)
func updateMatrix(mat [][]int) [][]int {
	dirs := [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	rowLen, colLen := len(mat), len(mat[0])
	queue := make([][2]int, 0)
	// 使用一個 queue 來存放數值為零的位置
	for row := 0; row < rowLen; row++ {
		for col := 0; col < colLen; col++ {
			if mat[row][col] == 0 {
				queue = append(queue, [2]int{row, col})
				continue
			}
			mat[row][col] = math.MaxInt
		}
	}

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		for _, dir := range dirs {
			row, col := cur[0]+dir[0], cur[1]+dir[1]
			// 檢查越界跟排除零的可能
			if row < 0 || row >= rowLen || col < 0 || col >= colLen || mat[row][col] <= mat[cur[0]][cur[1]] {
				continue
			}
			// 這時將剩下不為零的值開始丟進 queue 繼續累加
			queue = append(queue, [2]int{row, col})
			mat[row][col] = mat[cur[0]][cur[1]] + 1
		}
	}

	return mat
}

//leetcode submit region end(Prohibit modification and deletion)
