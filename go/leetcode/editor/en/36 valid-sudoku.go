package main

import "strconv"

//Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be
//validated according to the following rules:
//
//
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9
//without repetition.
//
//
// Note:
//
//
// A Sudoku board (partially filled) could be valid but is not necessarily
//solvable.
// Only the filled cells need to be validated according to the mentioned rules.
//
//
//
//
// Example 1:
//
//
//Input: board =
//[["5","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//Output: true
//
//
// Example 2:
//
//
//Input: board =
//[["8","3",".",".","7",".",".",".","."]
//,["6",".",".","1","9","5",".",".","."]
//,[".","9","8",".",".",".",".","6","."]
//,["8",".",".",".","6",".",".",".","3"]
//,["4",".",".","8",".","3",".",".","1"]
//,["7",".",".",".","2",".",".",".","6"]
//,[".","6",".",".",".",".","2","8","."]
//,[".",".",".","4","1","9",".",".","5"]
//,[".",".",".",".","8",".",".","7","9"]]
//Output: false
//Explanation: Same as Example 1, except with the 5 in the top left corner
//being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is
//invalid.
//
//
//
// Constraints:
//
//
// board.length == 9
// board[i].length == 9
// board[i][j] is a digit 1-9 or '.'.
//
//
// Related Topics Array Hash Table Matrix 👍 7623 👎 851

//leetcode submit region begin(Prohibit modification and deletion)
// https://leetcode.com/problems/valid-sudoku/solutions/1135904/0ms-simple-go-code/
// https://leetcode.com/problems/valid-sudoku/solutions/1214152/golang-solution-with-explanation-and-images-primes/
func isValidSudoku(board [][]byte) bool {
	rowMap := [9][9]bool{}
	colMap := [9][9]bool{}
	gridMap := [9][9]bool{}

	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			val, err := strconv.Atoi(string(board[row][col]))
			if err != nil {
				continue
			}
			// 將存入的數值 1 - 9 減一對應到 index
			val--
			// 利用 [0,1,2,3,4,5,6,7,8] 經過 / 3 後，會變成 [0,0,0,1,1,1,2,2,2]
			gridIdx := col/3 + (row/3)*3
			if rowMap[row][val] || colMap[col][val] || gridMap[gridIdx][val] {
				return false
			}

			rowMap[row][val] = true
			colMap[col][val] = true
			gridMap[gridIdx][val] = true
		}
	}

	return true
}

//leetcode submit region end(Prohibit modification and deletion)
