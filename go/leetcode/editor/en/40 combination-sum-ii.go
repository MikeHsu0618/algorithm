package main

import "sort"

//Given a collection of candidate numbers (candidates) and a target number (
//target), find all unique combinations in candidates where the candidate numbers sum
//to target.
//
// Each number in candidates may only be used once in the combination.
//
// Note: The solution set must not contain duplicate combinations.
//
//
// Example 1:
//
//
//Input: candidates = [10,1,2,7,6,1,5], target = 8
//Output:
//[
//[1,1,6],
//[1,2,5],
//[1,7],
//[2,6]
//]
//
//
// Example 2:
//
//
//Input: candidates = [2,5,2,1,2], target = 5
//Output:
//[
//[1,2,2],
//[5]
//]
//
//
//
// Constraints:
//
//
// 1 <= candidates.length <= 100
// 1 <= candidates[i] <= 50
// 1 <= target <= 30
//
//
// Related Topics Array Backtracking 👍 7444 👎 183

//leetcode submit region begin(Prohibit modification and deletion)
func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	sum := 0
	sort.Ints(candidates)
	var backtracing func(start int)
	backtracing = func(start int) {
		if sum > target {
			return
		}
		if sum == target {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(candidates); i++ {
			// 樹層去重
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}
			sum += candidates[i]
			path = append(path, candidates[i])
			backtracing(i + 1) // 去重
			sum -= candidates[i]
			path = path[:len(path)-1]
		}
	}
	backtracing(0)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
