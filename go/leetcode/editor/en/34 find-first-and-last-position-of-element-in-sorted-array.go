package main

//Given an array of integers nums sorted in non-decreasing order, find the
//starting and ending position of a given target value.
//
// If target is not found in the array, return [-1, -1].
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [5,7,7,8,8,10], target = 8
//Output: [3,4]
//
// Example 2:
// Input: nums = [5,7,7,8,8,10], target = 6
//Output: [-1,-1]
//
// Example 3:
// Input: nums = [], target = 0
//Output: [-1,-1]
//
//
// Constraints:
//
//
// 0 <= nums.length <= 10⁵
// -10⁹ <= nums[i] <= 10⁹
// nums is a non-decreasing array.
// -10⁹ <= target <= 10⁹
//
//
// Related Topics Array Binary Search 👍 15051 👎 363

//leetcode submit region begin(Prohibit modification and deletion)
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return []int{binarySearchLowBound(nums, target), binarySearchUpBound(nums, target)}
}

func binarySearchLowBound(nums []int, target int) int {
	res := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		res = mid
		r = mid - 1
	}

	return res
}

func binarySearchUpBound(nums []int, target int) int {
	res := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		res = mid
		l = mid + 1
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
