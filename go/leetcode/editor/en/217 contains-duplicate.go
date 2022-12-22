package main

import "sort"

//Given an integer array nums, return true if any value appears at least twice
//in the array, and return false if every element is distinct.
//
//
// Example 1:
// Input: nums = [1,2,3,1]
//Output: true
//
// Example 2:
// Input: nums = [1,2,3,4]
//Output: false
//
// Example 3:
// Input: nums = [1,1,1,3,3,4,3,2,4,2]
//Output: true
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
//
//
// Related Topics Array Hash Table Sorting 👍 7642 👎 1070

//leetcode submit region begin(Prohibit modification and deletion)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

func containsDuplicate1(nums []int) bool {
	m := make(map[int]bool, 0)

	for i := 0; i < len(nums); i++ {
		if v, ok := m[nums[i]]; ok {
			return v
		}
		m[nums[i]] = true
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
