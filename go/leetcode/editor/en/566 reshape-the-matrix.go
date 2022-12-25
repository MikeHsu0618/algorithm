package main

//In MATLAB, there is a handy function called reshape which can reshape an m x
//n matrix into a new one with a different size r x c keeping its original data.
//
// You are given an m x n matrix mat and two integers r and c representing the
//number of rows and the number of columns of the wanted reshaped matrix.
//
// The reshaped matrix should be filled with all the elements of the original
//matrix in the same row-traversing order as they were.
//
// If the reshape operation with given parameters is possible and legal, output
//the new reshaped matrix; Otherwise, output the original matrix.
//
//
// Example 1:
//
//
//Input: mat = [[1,2],[3,4]], r = 1, c = 4
//Output: [[1,2,3,4]]
//
//
// Example 2:
//
//
//Input: mat = [[1,2],[3,4]], r = 2, c = 4
//Output: [[1,2],[3,4]]
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat[i].length
// 1 <= m, n <= 100
// -1000 <= mat[i][j] <= 1000
// 1 <= r, c <= 300
//
//
// Related Topics Array Matrix Simulation 👍 2936 👎 325

//leetcode submit region begin(Prohibit modification and deletion)
func matrixReshape(mat [][]int, r int, c int) [][]int {
	// 檢查是否合法
	if r*c != len(mat)*len(mat[0]) {
		return mat
	}

	res := make([][]int, 0)
	path := make([]int, 0)
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			path = append(path, mat[i][j])
			if len(path) == c {
				res = append(res, append([]int{}, path...))
				path = make([]int, 0)
			}
		}
	}

	// 檢查是否合法
	// if len(res) != r || len(res[len(res) - 1]) != c {
	//     return mat
	// }

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
