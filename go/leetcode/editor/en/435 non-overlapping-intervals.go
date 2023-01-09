package main

import "sort"

//Given an array of intervals intervals where intervals[i] = [starti, endi],
//return the minimum number of intervals you need to remove to make the rest of the
//intervals non-overlapping.
//
//
// Example 1:
//
//
//Input: intervals = [[1,2],[2,3],[3,4],[1,3]]
//Output: 1
//Explanation: [1,3] can be removed and the rest of the intervals are non-
//overlapping.
//
//
// Example 2:
//
//
//Input: intervals = [[1,2],[1,2],[1,2]]
//Output: 2
//Explanation: You need to remove two [1,2] to make the rest of the intervals
//non-overlapping.
//
//
// Example 3:
//
//
//Input: intervals = [[1,2],[2,3]]
//Output: 0
//Explanation: You don't need to remove any of the intervals since they're
//already non-overlapping.
//
//
//
// Constraints:
//
//
// 1 <= intervals.length <= 10⁵
// intervals[i].length == 2
// -5 * 10⁴ <= starti < endi <= 5 * 10⁴
//
//
// Related Topics Array Dynamic Programming Greedy Sorting 👍 5364 👎 151

//leetcode submit region begin(Prohibit modification and deletion)
// 按照右边界排序，就要从左向右遍历，因为右边界越小越好，只要右边界越小，留给下一个区间的空间就越大，所以从左向右遍历，优先选右边界小的
// 按照左边界排序，就要从右向左遍历，因为左边界数值越大越好（越靠右），这样就给前一个区间的空间就越大，所以可以从右向左遍历
func eraseOverlapIntervals(intervals [][]int) int {
	// 使用右邊界排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	// 最後一定會至少剩一個
	count := 1
	for i := 1; i < len(intervals); i++ {
		// i 左邊界 ≥ i-1 右邊界，就不用切
		if intervals[i][0] >= intervals[i-1][1] {
			count++
			continue
		}
		// 更新右邊界
		intervals[i][1] = intervals[i-1][1]
	}

	return len(intervals) - count
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

//leetcode submit region end(Prohibit modification and deletion)
