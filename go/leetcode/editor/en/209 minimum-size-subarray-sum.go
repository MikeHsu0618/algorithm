package main

import "fmt"

//Given an array of positive integers nums and a positive integer target,
//return the minimal length of a subarray whose sum is greater than or equal to target.
//If there is no such subarray, return 0 instead.
//
//
// Example 1:
//
//
//Input: target = 7, nums = [2,3,1,2,4,3]
//Output: 2
//Explanation: The subarray [4,3] has the minimal length under the problem
//constraint.
//
//
// Example 2:
//
//
//Input: target = 4, nums = [1,4,4]
//Output: 1
//
//
// Example 3:
//
//
//Input: target = 11, nums = [1,1,1,1,1,1,1,1]
//Output: 0
//
//
//
// Constraints:
//
//
// 1 <= target <= 10⁹
// 1 <= nums.length <= 10⁵
// 1 <= nums[i] <= 10⁴
//
//
//
//Follow up: If you have figured out the
//O(n) solution, try coding another solution of which the time complexity is
//O(n log(n)).
//
// Related Topics Array Binary Search Sliding Window Prefix Sum 👍 8186 👎 226

//leetcode submit region begin(Prohibit modification and deletion)
// 使用 slice window 找到大於 target 的數值後，開始移動 start 指針
func minSubArrayLen(target int, nums []int) int {
	minLen := 0
	sum := 0
	start := 0
	for end := 0; end < len(nums); end++ {
		if nums[end] == target {
			return 1
		}
		sum += nums[end]
		for sum >= target {
			if minLen == 0 || end-start+1 < minLen {
				minLen = end - start + 1
			}
			sum -= nums[start]
			start++
		}
	}
	return minLen
}

// 解法一（暴力解）： Time limit exceeded
func minSubArrayLen1(target int, nums []int) int {
	min := 0
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			min = 1
			continue
		}
		count = nums[i]
		for j := i + 1; j < len(nums); j++ {
			count += nums[j]
			fmt.Println(count)
			if count < target {
				continue
			}
			if min == 0 || j-i+1 < min {
				min = j - i + 1
				break
			}
		}
	}
	return min
}

//leetcode submit region end(Prohibit modification and deletion)
