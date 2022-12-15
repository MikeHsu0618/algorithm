package main

import "sort"

//Given an array of intervals where intervals[i] = [starti, endi], merge all
//overlapping intervals, and return an array of the non-overlapping intervals that
//cover all the intervals in the input.
//
//
// Example 1:
//
//
//Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
//Output: [[1,6],[8,10],[15,18]]
//Explanation: Since intervals [1,3] and [2,6] overlap, merge them into [1,6].
//
//
// Example 2:
//
//
//Input: intervals = [[1,4],[4,5]]
//Output: [[1,5]]
//Explanation: Intervals [1,4] and [4,5] are considered overlapping.
//
//
//
// Constraints:
//
//
// 1 <= intervals.length <= 10⁴
// intervals[i].length == 2
// 0 <= starti <= endi <= 10⁴
//
//
// Related Topics Array Sorting 👍 17299 👎 613

//leetcode submit region begin(Prohibit modification and deletion)
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)

	for i := 0; i < len(intervals); {
		left, right := intervals[i][0], intervals[i][1]

		j := i + 1
		for j < len(intervals) && right >= intervals[j][0] {
			right = max(right, intervals[j][1])
			j++
		}

		res = append(res, []int{left, right})
		i = j
	}

	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
