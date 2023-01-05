package main

//Given an array nums of distinct integers, return all the possible
//permutations. You can return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
// Example 2:
// Input: nums = [0,1]
//Output: [[0,1],[1,0]]
//
// Example 3:
// Input: nums = [1]
//Output: [[1]]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 6
// -10 <= nums[i] <= 10
// All the integers of nums are unique.
//
//
// Related Topics Array Backtracking 👍 14190 👎 240

//leetcode submit region begin(Prohibit modification and deletion)
func permute(nums []int) [][]int {
	m := make(map[int]bool, 0)
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums); i++ {
			if m[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			m[nums[i]] = true
			backtracking(0)
			path = path[:len(path)-1]
			delete(m, nums[i])
		}
	}
	backtracking(0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
