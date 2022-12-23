package main

import (
	"fmt"
	"sort"
)

//You are given an m x n binary matrix mat of 1's (representing soldiers) and 0
//'s (representing civilians). The soldiers are positioned in front of the
//civilians. That is, all the 1's will appear to the left of all the 0's in each row.
//
// A row i is weaker than a row j if one of the following is true:
//
//
// The number of soldiers in row i is less than the number of soldiers in row j.
//
// Both rows have the same number of soldiers and i < j.
//
//
// Return the indices of the k weakest rows in the matrix ordered from weakest
//to strongest.
//
//
// Example 1:
//
//
//Input: mat =
//[[1,1,0,0,0],
// [1,1,1,1,0],
// [1,0,0,0,0],
// [1,1,0,0,0],
// [1,1,1,1,1]],
//k = 3
//Output: [2,0,3]
//Explanation:
//The number of soldiers in each row is:
//- Row 0: 2
//- Row 1: 4
//- Row 2: 1
//- Row 3: 2
//- Row 4: 5
//The rows ordered from weakest to strongest are [2,0,3,1,4].
//
//
// Example 2:
//
//
//Input: mat =
//[[1,0,0,0],
// [1,1,1,1],
// [1,0,0,0],
// [1,0,0,0]],
//k = 2
//Output: [0,2]
//Explanation:
//The number of soldiers in each row is:
//- Row 0: 1
//- Row 1: 4
//- Row 2: 1
//- Row 3: 1
//The rows ordered from weakest to strongest are [0,2,3,1].
//
//
//
// Constraints:
//
//
// m == mat.length
// n == mat[i].length
// 2 <= n, m <= 100
// 1 <= k <= m
// matrix[i][j] is either 0 or 1.
//
//
// Related Topics Array Binary Search Sorting Heap (Priority Queue) Matrix 👍 29
//71 👎 175

//leetcode submit region begin(Prohibit modification and deletion)
// 使用 binary search + sort
func kWeakestRows(mat [][]int, k int) []int {
	res := make([]int, k)
	items := make([]item, len(mat))

	for i := 0; i < len(mat); i++ {
		count := 0
		count = binarySearch(mat[i], 0)
		items[i] = item{i, count}
		fmt.Println(items)
	}

	sort.Slice(items, func(i, j int) bool {
		a, b := items[i], items[j]
		return a.count < b.count || (a.count == b.count && a.index < b.index)
	})

	for i := 0; i < k; i++ {
		res[i] = items[i].index
	}

	return res
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			r = mid - 1
			continue
		}
		l = mid + 1
	}
	return l
}

type item struct {
	index int
	count int
}

// 只使用 sort 解法 O(m*n + mlogm)
func kWeakestRows1(mat [][]int, k int) []int {
	res := make([]int, k)
	items := make([]item, len(mat))

	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[i]); j++ {
			count += mat[i][j]
		}
		items[i] = item{i, count}
	}

	sort.Slice(items, func(i, j int) bool {
		a, b := items[i], items[j]
		return a.count < b.count || (a.count == b.count && a.index < b.index)
	})

	for i := 0; i < k; i++ {
		res[i] = items[i].index
	}

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
