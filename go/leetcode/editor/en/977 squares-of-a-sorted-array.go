package main

//Given an integer array nums sorted in non-decreasing order, return an array
//of the squares of each number sorted in non-decreasing order.
//
//
// Example 1:
//
//
//Input: nums = [-4,-1,0,3,10]
//Output: [0,1,9,16,100]
//Explanation: After squaring, the array becomes [16,1,0,9,100].
//After sorting, it becomes [0,1,9,16,100].
//
//
// Example 2:
//
//
//Input: nums = [-7,-3,2,3,11]
//Output: [4,9,9,49,121]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁴
// -10⁴ <= nums[i] <= 10⁴
// nums is sorted in non-decreasing order.
//
//
//
//Follow up: Squaring each element and sorting the new array is very trivial,
//could you find an
//O(n) solution using a different approach?
//
// Related Topics Array Two Pointers Sorting 👍 6755 👎 170

//leetcode submit region begin(Prohibit modification and deletion)
// 解法一
// 雙指針：使用左側的負數平方與右側比較，大於右側則插入 res 最前側位置
func sortedSquares(nums []int) []int {
	var res []int
	left, right := 0, len(nums)-1
	for left <= right {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		if leftSquare > rightSquare {
			// 將 leftSquare 插入 res 最前端
			res = append([]int{leftSquare}, res...)
			left++
		} else {
			// 將 right 插入 res 最前端
			res = append([]int{rightSquare}, res...)
			right--
		}
	}
	return res
}

// 優化解法一：預先設定 res 的長度，省去重新分配的 runtime
func sortedSquares1(nums []int) []int {
	left, right, current := 0, len(nums)-1, len(nums)-1
	res := make([]int, len(nums))
	for left <= right {
		leftSquare, rightSquare := nums[left]*nums[left], nums[right]*nums[right]
		if leftSquare > rightSquare {
			res[current] = leftSquare
			left++
		} else {
			res[current] = rightSquare
			right--
		}
		current--
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
