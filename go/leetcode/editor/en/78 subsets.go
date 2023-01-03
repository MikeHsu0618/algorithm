package main

//Given an integer array nums of unique elements, return all possible subsets (
//the power set).
//
// The solution set must not contain duplicate subsets. Return the solution in
//any order.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3]
//Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
//
//
// Example 2:
//
//
//Input: nums = [0]
//Output: [[],[0]]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// All the numbers of nums are unique.
//
//
// Related Topics Array Backtracking Bit Manipulation 👍 13151 👎 186

//leetcode submit region begin(Prohibit modification and deletion)
func subsets(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backtracking func(start int)
	backtracking = func(start int) {
		// 在最一開始就已經將 [] 加進 res
		res = append(res, append([]int{}, path...))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}
	backtracking(0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
