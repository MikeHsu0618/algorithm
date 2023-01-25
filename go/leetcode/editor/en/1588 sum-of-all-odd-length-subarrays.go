package main

//Given an array of positive integers arr, return the sum of all possible odd-
//length subarrays of arr.
//
// A subarray is a contiguous subsequence of the array.
//
//
// Example 1:
//
//
//Input: arr = [1,4,2,5,3]
//Output: 58
//Explanation: The odd-length subarrays of arr and their sums are:
//[1] = 1
//[4] = 4
//[2] = 2
//[5] = 5
//[3] = 3
//[1,4,2] = 7
//[4,2,5] = 11
//[2,5,3] = 10
//[1,4,2,5,3] = 15
//If we add all these together we get 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58
//
//
// Example 2:
//
//
//Input: arr = [1,2]
//Output: 3
//Explanation: There are only 2 subarrays of odd length, [1] and [2]. Their sum
//is 3.
//
// Example 3:
//
//
//Input: arr = [10,11,12]
//Output: 66
//
//
//
// Constraints:
//
//
// 1 <= arr.length <= 100
// 1 <= arr[i] <= 1000
//
//
//
// Follow up:
//
// Could you solve this problem in O(n) time complexity?
//
// Related Topics Array Math Prefix Sum 👍 2937 👎 222

//leetcode submit region begin(Prohibit modification and deletion)
func sumOddLengthSubarrays(arr []int) int {
	//instead of calculating all subarray sum
	//we can think about how many times for each num in arr is calculated in the subarray
	//for each nums[i], to left, there are 0, 1, i elements to the left choices, this is (i+1) choices
	//to the right, there are 0, 1, n-1-i-1+1(n-i-1), total (n-i) choices
	//and ((i+1) * (n-i) + 1) / 2 choices are of odd length
	res := 0
	n := len(arr)
	for i, num := range arr {
		cnt := ((i+1)*(n-i) + 1) / 2
		res += num * cnt
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
