package main

//You are a professional robber planning to rob houses along a street. Each
//house has a certain amount of money stashed, the only constraint stopping you from
//robbing each of them is that adjacent houses have security systems connected and
//it will automatically contact the police if two adjacent houses were broken
//into on the same night.
//
// Given an integer array nums representing the amount of money of each house,
//return the maximum amount of money you can rob tonight without alerting the
//police.
//
//
// Example 1:
//
//
//Input: nums = [1,2,3,1]
//Output: 4
//Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
//Total amount you can rob = 1 + 3 = 4.
//
//
// Example 2:
//
//
//Input: nums = [2,7,9,3,1]
//Output: 12
//Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5
//(money = 1).
//Total amount you can rob = 2 + 9 + 1 = 12.
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 400
//
//
// Related Topics Array Dynamic Programming 👍 16475 👎 324

//leetcode submit region begin(Prohibit modification and deletion)
func rob(nums []int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	// 注意當前狀態跟先前狀態有一定的依賴關係
	m := make(map[int]int, 0)
	// i 等於考慮下標 i 的房屋，m[i] 等於可偷的總金額
	m[0] = nums[0]
	m[1] = max(nums[1], nums[0])

	for i := 2; i < len(nums); i++ {
		m[i] = max(m[i-2]+nums[i], m[i-1])
	}

	return m[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
