package main

//Write an efficient algorithm that searches for a value target in an m x n
//integer matrix matrix. This matrix has the following properties:
//
//
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the
//previous row.
//
//
//
// Example 1:
//
//
//Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
//Output: true
//
//
// Example 2:
//
//
//Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
//Output: false
//
//
//
// Constraints:
//
//
// m == matrix.length
// n == matrix[i].length
// 1 <= m, n <= 100
// -10⁴ <= matrix[i][j], target <= 10⁴
//
//
// Related Topics Array Binary Search Matrix 👍 10760 👎 323

//leetcode submit region begin(Prohibit modification and deletion)
func searchMatrix(matrix [][]int, target int) bool {
	rowLen, colLen := len(matrix), len(matrix[0])
	l, r := 0, rowLen*colLen-1

	for l <= r {
		mid := (l + r) / 2

		row, col := mid/colLen, mid%colLen
		if matrix[row][col] > target {
			r = mid - 1
			continue
		}

		if matrix[row][col] < target {
			l = mid + 1
			continue
		}
		return true
	}

	return false
}

func searchMatrix1(matrix [][]int, target int) bool {
	// 需要從 row 由大到小檢查下來
	r, c := len(matrix)-1, 0

	for r >= 0 && c < len(matrix[0]) {
		// 移動到開頭小於等於 target 的 row
		if matrix[r][c] > target {
			r--
			continue
		}
		// 遍歷 column
		for c < len(matrix[0]) {
			if matrix[r][c] != target {
				c++
				continue
			}
			return true
		}
		break
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
