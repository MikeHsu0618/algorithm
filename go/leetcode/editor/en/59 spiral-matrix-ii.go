package main

//Given a positive integer n, generate an n x n matrix filled with elements
//from 1 to n² in spiral order.
//
//
// Example 1:
//
//
//Input: n = 3
//Output: [[1,2,3],[8,9,4],[7,6,5]]
//
//
// Example 2:
//
//
//Input: n = 1
//Output: [[1]]
//
//
//
// Constraints:
//
//
// 1 <= n <= 20
//
//
// Related Topics Array Matrix Simulation 👍 4338 👎 199

//leetcode submit region begin(Prohibit modification and deletion)
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, n)
	}
	left, right := 0, n-1
	top, bottom := 0, n-1
	num := 1
	tar := n * n
	for num <= tar {
		// left -> right 左閉右開
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++
		// top -> bottom 上閉下開
		for i := top; i <= bottom; i++ {
			res[i][bottom] = num
			num++
		}
		right--
		// right -> left 右閉左開
		for i := right; i >= left; i-- {
			res[bottom][i] = num
			num++
		}
		bottom--
		// bottom -> top 下閉上開
		for i := bottom; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
