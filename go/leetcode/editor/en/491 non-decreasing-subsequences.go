package main

//Given an integer array nums, return all the different possible non-decreasing
//subsequences of the given array with at least two elements. You may return the
//answer in any order.
//
//
// Example 1:
//
//
//Input: nums = [4,6,7,7]
//Output: [[4,6],[4,6,7],[4,6,7,7],[4,7],[4,7,7],[6,7],[6,7,7],[7,7]]
//
//
// Example 2:
//
//
//Input: nums = [4,4,3,2,1]
//Output: [[4,4]]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 15
// -100 <= nums[i] <= 100
//
//
// Related Topics Array Hash Table Backtracking Bit Manipulation 👍 1804 👎 161

//leetcode submit region begin(Prohibit modification and deletion)
func findSubsequences(nums []int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtracking func(start int)
	backtracking = func(start int) {
		// 注意不要 return，讓 path 繼續遍歷
		if len(path) > 1 {
			res = append(res, append([]int{}, path...))
		}

		// 需要對同一層的數值去重，故需要在這裡使用 hash table
		m := make(map[int]bool, 0)
		for i := start; i < len(nums); i++ {
			// 去重
			if m[nums[i]] {
				continue
			}
			// 檢查是否升序
			if len(path) != 0 && nums[i] < path[len(path)-1] {
				continue
			}
			path = append(path, nums[i])
			m[nums[i]] = true
			backtracking(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtracking(0)

	return res
}

//leetcode submit region end(Prohibit modification and deletion)
