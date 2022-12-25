package main

//There is an integer array nums sorted in ascending order (with distinct
//values).
//
// Prior to being passed to your function, nums is possibly rotated at an
//unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k]
//, nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For
//example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0
//,1,2].
//
// Given the array nums after the possible rotation and an integer target,
//return the index of target if it is in nums, or -1 if it is not in nums.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
// Input: nums = [4,5,6,7,0,1,2], target = 0
//Output: 4
//
// Example 2:
// Input: nums = [4,5,6,7,0,1,2], target = 3
//Output: -1
//
// Example 3:
// Input: nums = [1], target = 0
//Output: -1
//
//
// Constraints:
//
//
// 1 <= nums.length <= 5000
// -10⁴ <= nums[i] <= 10⁴
// All values of nums are unique.
// nums is an ascending array that is possibly rotated.
// -10⁴ <= target <= 10⁴
//
//
// Related Topics Array Binary Search 👍 18402 👎 1094

//leetcode submit region begin(Prohibit modification and deletion)
// 使用二分搜尋法，並判斷 target 位於哪一側，直到找到為止
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2

		if nums[mid] == target {
			return mid
		}

		// 關鍵在於確認哪一邊為 sorted side
		if nums[mid] < nums[r] {
			// 如果 target 不再這個範圍就把這個 sorted side 捨棄
			if target > nums[mid] && nums[r] >= target {
				l = mid + 1
				continue
			}
			// 如果超出邊界，左邊的數值一定比右邊大
			r = mid - 1
			continue
		}

		if nums[l] <= nums[mid] {
			if nums[mid] > target && target >= nums[l] {
				r = mid - 1
				continue
			}
			l = mid + 1
			continue
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
