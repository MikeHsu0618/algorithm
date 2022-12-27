package main

//Given an array, rotate the array to the right by k steps, where k is non-
//negative.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,4,5,6,7], k = 3
//Output: [5,6,7,1,2,3,4]
//Explanation:
//rotate 1 steps to the right: [7,1,2,3,4,5,6]
//rotate 2 steps to the right: [6,7,1,2,3,4,5]
//rotate 3 steps to the right: [5,6,7,1,2,3,4]
//
//
// Example 2:
//
//
//Input: nums = [-1,-100,3,99], k = 2
//Output: [3,99,-1,-100]
//Explanation:
//rotate 1 steps to the right: [99,-1,-100,3]
//rotate 2 steps to the right: [3,99,-1,-100]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -2³¹ <= nums[i] <= 2³¹ - 1
// 0 <= k <= 10⁵
//
//
//
// Follow up:
//
//
// Try to come up with as many solutions as you can. There are at least three
//different ways to solve this problem.
// Could you do it in-place with O(1) extra space?
//
//
// Related Topics Array Math Two Pointers 👍 12488 👎 1478

//leetcode submit region begin(Prohibit modification and deletion)
func rotate(nums []int, k int) {
	n := len(nums)
	// k 有可能大於 n
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 利用 最後面 k 個數值會與剩餘的前面互換，注意 copy slice 的指針
func rotate1(nums []int, k int) {
	n := len(nums)
	k %= n
	newNums := make([]int, len(nums))
	copy(newNums[:k], nums[len(nums)-k:])
	copy(newNums[k:], nums[:len(nums)-k])
	copy(nums, newNums)
}

//leetcode submit region end(Prohibit modification and deletion)
