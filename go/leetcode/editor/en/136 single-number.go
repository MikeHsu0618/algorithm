package main

//Given a non-empty array of integers nums, every element appears twice except
//for one. Find that single one.
//
// You must implement a solution with a linear runtime complexity and use only
//constant extra space.
//
//
// Example 1:
// Input: nums = [2,2,1]
//Output: 1
//
// Example 2:
// Input: nums = [4,1,2,1,2]
//Output: 4
//
// Example 3:
// Input: nums = [1]
//Output: 1
//
//
// Constraints:
//
//
// 1 <= nums.length <= 3 * 10⁴
// -3 * 10⁴ <= nums[i] <= 3 * 10⁴
// Each element in the array appears twice except for one element which appears
//only once.
//
//
// Related Topics Array Bit Manipulation 👍 12561 👎 479

//leetcode submit region begin(Prohibit modification and deletion)
// 使用 ^ XOR 運算子，可以發現只要兩次 XOR 皆可以還原原本的數值
// nums = [2,2,1]
// fmt.Printf("前: %b, %b , %v, %v\n", result, num, result, num)
// result ^= num
// fmt.Printf("後: %b, %b , %v, %v\n", result, num, result, num)
// 前: 0, 10 , 0, 2
// 後: 10, 10 , 2, 2
// 前: 10, 10 , 2, 2
// 後: 0, 10 , 0, 2
// 前: 0, 1 , 0, 1
// 後: 1, 1 , 1, 1
func singleNumber(nums []int) int {
	var result int
	for _, num := range nums {
		result ^= num
	}
	return result
}

// hash table
func singleNumber1(nums []int) int {
	m := make(map[int]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}

	for i, v := range m {
		if v == 1 {
			return i
		}
	}

	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
