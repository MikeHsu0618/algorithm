package main

import "sort"

//Given an integer array nums, return all the triplets [nums[i], nums[j], nums[
//k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
//
// Example 1:
//
//
//Input: nums = [-1,0,1,2,-1,-4]
//Output: [[-1,-1,2],[-1,0,1]]
//Explanation:
//nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
//nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
//nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
//The distinct triplets are [-1,0,1] and [-1,-1,2].
//Notice that the order of the output and the order of the triplets does not
//matter.
//
//
// Example 2:
//
//
//Input: nums = [0,1,1]
//Output: []
//Explanation: The only possible triplet does not sum up to 0.
//
//
// Example 3:
//
//
//Input: nums = [0,0,0]
//Output: [[0,0,0]]
//Explanation: The only possible triplet sums up to 0.
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
//
// Related Topics Array Two Pointers Sorting 👍 22270 👎 2039

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	// 遍歷數組，將問題視為拆解成當前數值和其他數值相加為零
	// 使用雙指針來找尋排序後的剩餘數組中相加為零的三個數值
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		// 刪除相同數字的值(當前數值可能與上一個數值相同，將會導致重複記錄)
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}
		begin := i + 1
		end := len(nums) - 1
		for begin < end {
			sum := nums[i] + nums[begin] + nums[end]
			// 如果數值過大，則 end 往 begin 移動（因為數組已經過排序）
			if sum > 0 {
				end--
				continue
			}
			// 反之，begin 往 end 移動
			if sum < 0 {
				begin++
				continue
			}
			result = append(result, []int{nums[i], nums[begin], nums[end]})
			// 處理如果 begin 和 end 下一位數值與前一位重複
			for begin < end && nums[end] == nums[end-1] {
				end--
			}
			for begin < end && nums[begin] == nums[begin+1] {
				begin++
			}
			end--
			begin++
		}
	}
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
