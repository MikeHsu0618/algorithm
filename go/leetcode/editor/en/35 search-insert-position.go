package main

//Given a sorted array of distinct integers and a target value, return the
//index if the target is found. If not, return the index where it would be if it were
//inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
//
//
//Input: nums = [1,3,5,6], target = 5
//Output: 2
//
//
// Example 2:
//
//
//Input: nums = [1,3,5,6], target = 2
//Output: 1
//
//
// Example 3:
//
//
//Input: nums = [1,3,5,6], target = 7
//Output: 4
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums contains distinct values sorted in ascending order.
// -10⁴ <= target <= 10⁴
//
//
// Related Topics Array Binary Search 👍 10849 👎 514

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
			continue
		}
		if target <= nums[mid] {
			r = mid - 1
			continue
		}
	}
	return l
}

func searchInsert1(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if target > nums[mid] {
			l = mid + 1
			continue
		}
		if target < nums[mid] {
			r = mid - 1
			continue
		}
		return mid
	}

	return l
}

//leetcode submit region end(Prohibit modification and deletion)
