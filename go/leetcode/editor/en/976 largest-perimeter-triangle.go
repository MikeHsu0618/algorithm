package main

import "sort"

//Given an integer array nums, return the largest perimeter of a triangle with
//a non-zero area, formed from three of these lengths. If it is impossible to form
//any triangle of a non-zero area, return 0.
//
//
// Example 1:
//
//
//Input: nums = [2,1,2]
//Output: 5
//Explanation: You can form a triangle with three side lengths: 1, 2, and 2.
//
//
// Example 2:
//
//
//Input: nums = [1,2,1,10]
//Output: 0
//Explanation:
//You cannot use the side lengths 1, 1, and 2 to form a triangle.
//You cannot use the side lengths 1, 1, and 10 to form a triangle.
//You cannot use the side lengths 1, 2, and 10 to form a triangle.
//As we cannot use any three side lengths to form a triangle of non-zero area,
//we return 0.
//
//
//
// Constraints:
//
//
// 3 <= nums.length <= 10⁴
// 1 <= nums[i] <= 10⁶
//
//
// Related Topics Array Math Greedy Sorting 👍 2542 👎 358

//leetcode submit region begin(Prohibit modification and deletion)
func largestPerimeter(nums []int) int {
	sort.Ints(nums)

	right := len(nums) - 1

	for right >= 2 {
		if nums[right-2] > nums[right-1]-nums[right] && nums[right-2] > nums[right]-nums[right-1] {
			return nums[right-2] + nums[right-1] + nums[right]
		}

		right--
	}

	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
