package main

//Given an integer array nums, find the subarray which has the largest sum and
//return its sum.
//
//
// Example 1:
//
//
//Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
//Output: 6
//Explanation: [4,-1,2,1] has the largest sum = 6.
//
//
// Example 2:
//
//
//Input: nums = [1]
//Output: 1
//
//
// Example 3:
//
//
//Input: nums = [5,4,-1,7,8]
//Output: 23
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
//
//
//
// Follow up: If you have figured out the O(n) solution, try coding another
//solution using the divide and conquer approach, which is more subtle.
//
// Related Topics Array Divide and Conquer Dynamic Programming 👍 26959 👎 1206

//leetcode submit region begin(Prohibit modification and deletion)
// 最大的子連集頭尾一定是正數
// 先找到那些頭為正整數的 key 值
func maxSubArray(nums []int) int {
	max, sum := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}

		if max < sum {
			max = sum
		}
	}

	return max
}

//leetcode submit region end(Prohibit modification and deletion)
