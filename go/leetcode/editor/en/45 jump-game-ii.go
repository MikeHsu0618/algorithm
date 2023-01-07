package main

//You are given a 0-indexed array of integers nums of length n. You are
//initially positioned at nums[0].
//
// Each element nums[i] represents the maximum length of a forward jump from
//index i. In other words, if you are at nums[i], you can jump to any nums[i + j]
//where:
//
//
// 0 <= j <= nums[i] and
// i + j < n
//
//
// Return the minimum number of jumps to reach nums[n - 1]. The test cases are
//generated such that you can reach nums[n - 1].
//
//
// Example 1:
//
//
//Input: nums = [2,3,1,1,4]
//Output: 2
//Explanation: The minimum number of jumps to reach the last index is 2. Jump 1
//step from index 0 to 1, then 3 steps to the last index.
//
//
// Example 2:
//
//
//Input: nums = [2,3,0,1,4]
//Output: 2
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// 0 <= nums[i] <= 1000
//
//
// Related Topics Array Dynamic Programming Greedy 👍 10461 👎 367

//leetcode submit region begin(Prohibit modification and deletion)
// 找到最大覆蓋範圍邊界為 len(nums) - 2 的最小步數（題目預設一定可以走到 len(nums) - 2, 所以不用擔心走不到）
func jump(nums []int) int {
	count := 0
	cover, nextCover := 0, 0
	// len(nums) - 2 為關鍵
	for i := 0; i < len(nums)-1; i++ {
		nextCover = max(nextCover, nums[i]+i)
		if i == cover {
			cover = nextCover
			count++
		}
	}

	return count
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
