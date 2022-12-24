package main

import "math"

//Given a non-negative integer c, decide whether there're two integers a and b
//such that a² + b² = c.
//
//
// Example 1:
//
//
//Input: c = 5
//Output: true
//Explanation: 1 * 1 + 2 * 2 = 5
//
//
// Example 2:
//
//
//Input: c = 3
//Output: false
//
//
//
// Constraints:
//
//
// 0 <= c <= 2³¹ - 1
//
//
// Related Topics Math Two Pointers Binary Search 👍 1893 👎 503

//leetcode submit region begin(Prohibit modification and deletion)
// i, j 一定 ≤ 根號 c, 並用 float64 防止最大值邊界
func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j

		if sum > c {
			j--
			continue
		}

		if sum < c {
			i++
			continue
		}

		return true
	}

	return false
}

// 沒有處理 c 的最大值，會超過 int 範圍
func judgeSquareSum1(c int) bool {
	// 遍歷 0 - c 使用 binarySearch 尋早符合值
	for i := 0; i <= c; i++ {
		target := c - i*i
		if target < 0 {
			continue
		}
		if !binarySearch(i, c, target) {
			continue
		}
		return true
	}
	return false
}

func binarySearch(l, r, target int) bool {
	for l <= r {
		mid := (l + r) / 2
		if mid*mid > target {
			r = mid - 1
			continue
		}

		if mid*mid < target {
			l = mid + 1
			continue
		}

		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
