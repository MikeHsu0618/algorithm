package main

import "sort"

//There are some spherical balloons taped onto a flat wall that represents the
//XY-plane. The balloons are represented as a 2D integer array points where points[
//i] = [xstart, xend] denotes a balloon whose horizontal diameter stretches
//between xstart and xend. You do not know the exact y-coordinates of the balloons.
//
// Arrows can be shot up directly vertically (in the positive y-direction) from
//different points along the x-axis. A balloon with xstart and xend is burst by
//an arrow shot at x if xstart <= x <= xend. There is no limit to the number of
//arrows that can be shot. A shot arrow keeps traveling up infinitely, bursting any
//balloons in its path.
//
// Given the array points, return the minimum number of arrows that must be
//shot to burst all balloons.
//
//
// Example 1:
//
//
//Input: points = [[10,16],[2,8],[1,6],[7,12]]
//Output: 2
//Explanation: The balloons can be burst by 2 arrows:
//- Shoot an arrow at x = 6, bursting the balloons [2,8] and [1,6].
//- Shoot an arrow at x = 11, bursting the balloons [10,16] and [7,12].
//
//
// Example 2:
//
//
//Input: points = [[1,2],[3,4],[5,6],[7,8]]
//Output: 4
//Explanation: One arrow needs to be shot for each balloon for a total of 4
//arrows.
//
//
// Example 3:
//
//
//Input: points = [[1,2],[2,3],[3,4],[4,5]]
//Output: 2
//Explanation: The balloons can be burst by 2 arrows:
//- Shoot an arrow at x = 2, bursting the balloons [1,2] and [2,3].
//- Shoot an arrow at x = 4, bursting the balloons [3,4] and [4,5].
//
//
//
// Constraints:
//
//
// 1 <= points.length <= 10⁵
// points[i].length == 2
// -2³¹ <= xstart < xend <= 2³¹ - 1
//
//
// Related Topics Array Greedy Sorting 👍 5746 👎 159

//leetcode submit region begin(Prohibit modification and deletion)
func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	count := 1
	for i := 1; i < len(points); i++ {
		// 如果前一位的右邊界小於後一位的左邊界一定不符合
		if points[i-1][1] < points[i][0] {
			count++
			continue
		}
		// 更新重疊氣球的最小邊界，覆蓋該位置的值，留到下一步
		points[i][1] = min(points[i][1], points[i-1][1])
	}

	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
