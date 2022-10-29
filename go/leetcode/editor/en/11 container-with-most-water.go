package main

//You are given an integer array height of length n. There are n vertical lines
//drawn such that the two endpoints of the iᵗʰ line are (i, 0) and (i, height[i]).
//
//
// Find two lines that together with the x-axis form a container, such that the
//container contains the most water.
//
// Return the maximum amount of water a container can store.
//
// Notice that you may not slant the container.
//
//
// Example 1:
//
//
//Input: height = [1,8,6,2,5,4,8,3,7]
//Output: 49
//Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,
//3,7]. In this case, the max area of water (blue section) the container can
//contain is 49.
//
//
// Example 2:
//
//
//Input: height = [1,1]
//Output: 1
//
//
//
// Constraints:
//
//
// n == height.length
// 2 <= n <= 10⁵
// 0 <= height[i] <= 10⁴
//
//
// Related Topics Array Two Pointers Greedy 👍 20965 👎 1110

// 附上不會通過的暴力解
//func maxArea(height []int) int {
//	if len(height) == 2 {
//		if height[0] < height[1] {
//			return height[0]
//		}
//		return height[1]
//	}
//	max := 0
//	for i:= 0; i < len(height)-1; i++ {
//		// 暴力遍歷所有可能，得出最大值
//		for j:=i+1; j<len(height);j++ {
//			h := height[i]
//			if height[i] > height[j] {
//				h = height[j]
//			}
//			if h*(j-i) > max {
//				max = h*(j-i)
//			}
//		}
//	}
//	return max
//}

//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	start := 0
	end := len(height) - 1
	max := 0
	// 使用雙指針比較每次起點終點的長短
	// 使用較短者的高以及兩者距離做為每次最大值
	// 並將較短者往較長端做移動
	for start < end {
		// 最少都會有一
		current := 0
		if height[start] < height[end] {
			current = height[start] * (end - start)
			start++
		} else {
			current = height[end] * (end - start)
			end--
		}
		if current > max {
			max = current
		}
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
