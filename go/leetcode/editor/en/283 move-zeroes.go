package main

//Given an integer array nums, move all 0's to the end of it while maintaining
//the relative order of the non-zero elements.
//
// Note that you must do this in-place without making a copy of the array.
//
//
// Example 1:
// Input: nums = [0,1,0,3,12]
//Output: [1,3,12,0,0]
//
// Example 2:
// Input: nums = [0]
//Output: [0]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -2³¹ <= nums[i] <= 2³¹ - 1
//
//
//
//Follow up: Could you minimize the total number of operations done?
//
// Related Topics Array Two Pointers 👍 12158 👎 303

//leetcode submit region begin(Prohibit modification and deletion)
// 一個簡單的 loop
func moveZeroes(nums []int) {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
}

// 快慢指針從頭開始遍歷
func moveZeroes1(nums []int) {
	slow, fast := 0, 0

	for fast < len(nums) {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}

	for slow < len(nums) {
		nums[slow] = 0
		slow++
	}
}

//leetcode submit region end(Prohibit modification and deletion)
