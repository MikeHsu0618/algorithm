package main

import (
	"math"
	"sort"
)

//Given an integer array nums and an integer k, modify the array in the
//following way:
//
//
// choose an index i and replace nums[i] with -nums[i].
//
//
// You should apply this process exactly k times. You may choose the same index
//i multiple times.
//
// Return the largest possible sum of the array after modifying it in this way.
//
//
//
// Example 1:
//
//
//Input: nums = [4,2,3], k = 1
//Output: 5
//Explanation: Choose index 1 and nums becomes [4,-2,3].
//
//
// Example 2:
//
//
//Input: nums = [3,-1,0,2], k = 3
//Output: 6
//Explanation: Choose indices (1, 2, 2) and nums becomes [3,1,0,2].
//
//
// Example 3:
//
//
//Input: nums = [2,-3,-1,5,-4], k = 2
//Output: 13
//Explanation: Choose indices (1, 4) and nums becomes [2,3,-1,5,4].
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -100 <= nums[i] <= 100
// 1 <= k <= 10⁴
//
//
// Related Topics Array Greedy Sorting 👍 1220 👎 94

//leetcode submit region begin(Prohibit modification and deletion)
// 第一步：將數字組按絕對值大的小從大到小順序，注意要按照絕對值的大的小
// 第二步：從前向後遍歷，遇到負數將其變化為正數，同時K--
// 第三步：如果K還大於0，那麼反反复復改變數值最小的元素，將K用完
// 第四步：求和
func largestSumAfterKNegations(nums []int, k int) int {
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})

	for i := 0; i < len(nums) && k > 0; i++ {
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}

	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return result
}

// 不使用貪心法
func largestSumAfterKNegations1(nums []int, k int) int {
	// 排序後，盡量將負數變為正數，如數組都為正數且 k 為奇數，將最小的正數變為負數
	sort.Ints(nums)
	minNum := math.MaxInt
	index := 0
	for i := 0; i < len(nums) && k > 0; i++ {
		if ad(nums[i]) < minNum {
			minNum = ad(nums[i])
			index = i
		}

		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}

		if nums[i] == 0 {
			break
		}
	}

	for k > 0 {
		nums[index] = -nums[index]
		k--
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
	}

	return count
}

func ad(a int) int {
	if a > 0 {
		return a
	}

	return -a
}

//leetcode submit region end(Prohibit modification and deletion)
