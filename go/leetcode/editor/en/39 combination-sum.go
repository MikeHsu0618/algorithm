package main

//Given an array of distinct integers candidates and a target integer target,
//return a list of all unique combinations of candidates where the chosen numbers
//sum to target. You may return the combinations in any order.
//
// The same number may be chosen from candidates an unlimited number of times.
//Two combinations are unique if the frequency of at least one of the chosen
//numbers is different.
//
// The test cases are generated such that the number of unique combinations
//that sum up to target is less than 150 combinations for the given input.
//
//
// Example 1:
//
//
//Input: candidates = [2,3,6,7], target = 7
//Output: [[2,2,3],[7]]
//Explanation:
//2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple
//times.
//7 is a candidate, and 7 = 7.
//These are the only two combinations.
//
//
// Example 2:
//
//
//Input: candidates = [2,3,5], target = 8
//Output: [[2,2,2,2],[2,3,3],[3,5]]
//
//
// Example 3:
//
//
//Input: candidates = [2], target = 1
//Output: []
//
//
//
// Constraints:
//
//
// 1 <= candidates.length <= 30
// 2 <= candidates[i] <= 40
// All elements of candidates are distinct.
// 1 <= target <= 40
//
//
// Related Topics Array Backtracking 👍 14007 👎 285

//leetcode submit region begin(Prohibit modification and deletion)
// 回溯三部曲：遞迴參數、終止條件、單層判斷邏輯
// 優化：剪枝（提早退出）
func combinationSum(candidates []int, target int) [][]int {
	var track []int
	var res [][]int
	backtracking(0, 0, target, candidates, track, &res)
	return res
}

func backtracking(startIndex, sum, target int, candidates []int, track []int, res *[][]int) {
	// 終止條件
	if sum == target {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}
	// 提早退出
	if sum > target {
		return
	}
	// 開始回溯
	for i := startIndex; i < len(candidates); i++ {
		// 更新路徑集合和 sum
		track = append(track, candidates[i])
		sum += candidates[i]
		// 遞迴
		// startIndex 沒有傳入 i+1 是因為允許相同本身數值重複
		backtracking(i, sum, target, candidates, track, res)
		// 回溯
		track = track[:len(track)-1]
		sum -= candidates[i]
	}
}

//leetcode submit region end(Prohibit modification and deletion)
