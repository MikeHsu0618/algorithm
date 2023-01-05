package main

import "sort"

//Given a collection of numbers, nums, that might contain duplicates, return
//all possible unique permutations in any order.
//
//
// Example 1:
//
//
//Input: nums = [1,1,2]
//Output:
//[[1,1,2],
// [1,2,1],
// [2,1,1]]
//
//
// Example 2:
//
//
//Input: nums = [1,2,3]
//Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 8
// -10 <= nums[i] <= 10
//
//
// Related Topics Array Backtracking 👍 6911 👎 124

//leetcode submit region begin(Prohibit modification and deletion)
func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	path := make([]int, 0)
	m := make(map[int]bool, len(nums))
	var backtracking func(start int)
	backtracking = func(start int) {
		if len(path) == len(nums) {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i < len(nums); i++ {
			// m[i - 1] == true，说明同一树枝nums[i - 1]使用过
			// m[i - 1] == false，说明同一树层nums[i - 1]使用过
			// 如果同一树层nums[i - 1]使用过则直接跳过
			if i != start && nums[i] == nums[i-1] && m[i-1] {
				continue
			}
			if m[i] {
				continue
			}
			path = append(path, nums[i])
			m[i] = true
			backtracking(0)
			path = path[:len(path)-1]
			m[i] = false
		}
	}
	backtracking(0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
