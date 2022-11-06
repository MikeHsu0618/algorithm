package main

//Given an array of integers nums which is sorted in ascending order, and an
//integer target, write a function to search target in nums. If target exists, then
//return its index. Otherwise, return -1.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
// Example 1:
//
//
//Input: nums = [-1,0,3,5,9,12], target = 9
//Output: 4
//Explanation: 9 exists in nums and its index is 4
//
//
// Example 2:
//
//
//Input: nums = [-1,0,3,5,9,12], target = 2
//Output: -1
//Explanation: 2 does not exist in nums so return -1
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ < nums[i], target < 10⁴
// All the integers in nums are unique.
// nums is sorted in ascending order.
//
//
// Related Topics Array Binary Search 👍 7126 👎 155

//leetcode submit region begin(Prohibit modification and deletion)
// binary search works on a sorted array
func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		// int 的除法將會自動省略小數點
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		// 判斷 target 再 mid 的左側還是右側
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
