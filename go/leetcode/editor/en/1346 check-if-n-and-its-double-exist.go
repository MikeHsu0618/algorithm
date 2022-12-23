package main

import "sort"

//Given an array arr of integers, check if there exist two indices i and j such
//that :
//
//
// i != j
// 0 <= i, j < arr.length
// arr[i] == 2 * arr[j]
//
//
//
// Example 1:
//
//
//Input: arr = [10,2,5,3]
//Output: true
//Explanation: For i = 0 and j = 2, arr[i] == 10 == 2 * 5 == 2 * arr[j]
//
//
// Example 2:
//
//
//Input: arr = [3,1,7,11]
//Output: false
//Explanation: There is no i and j that satisfy the conditions.
//
//
//
// Constraints:
//
//
// 2 <= arr.length <= 500
// -10³ <= arr[i] <= 10³
//
//
// Related Topics Array Hash Table Two Pointers Binary Search Sorting 👍 1378 👎
// 151

//leetcode submit region begin(Prohibit modification and deletion)
func checkIfExist(arr []int) bool {
	sort.Ints(arr)
	for i := 0; i < len(arr); i++ {
		tmp := make([]int, len(arr))
		copy(tmp, arr)
		nums := append(tmp[:i], tmp[i+1:]...)
		if binarySearch(nums, 2*arr[i]) == false {
			continue
		}
		return true
	}

	return false
}

func binarySearch(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] > target {
			r = mid - 1
			continue
		}

		if nums[mid] < target {
			l = mid + 1
			continue
		}
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
