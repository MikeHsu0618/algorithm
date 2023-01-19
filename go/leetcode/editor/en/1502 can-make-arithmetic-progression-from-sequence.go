package main

import "sort"

//A sequence of numbers is called an arithmetic progression if the difference
//between any two consecutive elements is the same.
//
// Given an array of numbers arr, return true if the array can be rearranged to
//form an arithmetic progression. Otherwise, return false.
//
//
// Example 1:
//
//
//Input: arr = [3,5,1]
//Output: true
//Explanation: We can reorder the elements as [1,3,5] or [5,3,1] with
//differences 2 and -2 respectively, between each consecutive elements.
//
//
// Example 2:
//
//
//Input: arr = [1,2,4]
//Output: false
//Explanation: There is no way to reorder the elements to obtain an arithmetic
//progression.
//
//
//
// Constraints:
//
//
// 2 <= arr.length <= 1000
// -10⁶ <= arr[i] <= 10⁶
//
//
// Related Topics Array Sorting 👍 954 👎 58

//leetcode submit region begin(Prohibit modification and deletion)
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := abs(arr[0] - arr[1])
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) != diff {
			return false
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

//leetcode submit region end(Prohibit modification and deletion)
