package main

import "sort"

//Given an array nums of n integers, return an array of all the unique
//quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
//
//
// 0 <= a, b, c, d < n
// a, b, c, and d are distinct.
// nums[a] + nums[b] + nums[c] + nums[d] == target
//
//
// You may return the answer in any order.
//
//
// Example 1:
//
//
//Input: nums = [1,0,-1,0,-2,2], target = 0
//Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
//
//
// Example 2:
//
//
//Input: nums = [2,2,2,2,2], target = 8
//Output: [[2,2,2,2]]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 200
// -10⁹ <= nums[i] <= 10⁹
// -10⁹ <= target <= 10⁹
//
//
// Related Topics Array Two Pointers Sorting 👍 8081 👎 940

//leetcode submit region begin(Prohibit modification and deletion)

// 解法：基本上就是 3sum 再包一個迴圈
func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums)-3; i++ {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		others := target - nums[i]
		for j := i + 1; j < len(nums)-2; j++ {
			if j != i+1 && nums[j] == nums[j-1] {
				continue
			}
			begin, end := j+1, len(nums)-1
			for begin < end {
				sum := nums[j] + nums[begin] + nums[end]
				if sum < others {
					for begin < end && nums[begin] == nums[begin+1] {
						begin++
					}
					begin++
					continue
				}
				if sum > others {
					for begin < end && nums[end] == nums[end-1] {
						end--
					}
					end--
					continue
				}
				res = append(res, []int{nums[i], nums[j], nums[begin], nums[end]})
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
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
