package main

//There are n children standing in a line. Each child is assigned a rating
//value given in the integer array ratings.
//
// You are giving candies to these children subjected to the following
//requirements:
//
//
// Each child must have at least one candy.
// Children with a higher rating get more candies than their neighbors.
//
//
// Return the minimum number of candies you need to have to distribute the
//candies to the children.
//
//
// Example 1:
//
//
//Input: ratings = [1,0,2]
//Output: 5
//Explanation: You can allocate to the first, second and third child with 2, 1,
//2 candies respectively.
//
//
// Example 2:
//
//
//Input: ratings = [1,2,2]
//Output: 4
//Explanation: You can allocate to the first, second and third child with 1, 2,
//1 candies respectively.
//The third child gets 1 candy because it satisfies the above two conditions.
//
//
//
// Constraints:
//
//
// n == ratings.length
// 1 <= n <= 2 * 10⁴
// 0 <= ratings[i] <= 2 * 10⁴
//
//
// Related Topics Array Greedy 👍 5047 👎 333

//leetcode submit region begin(Prohibit modification and deletion)
// 貪心法：分兩次比較評分
func candy(ratings []int) int {
	res := 0
	counts := make([]int, len(ratings))

	// 初始化
	for i := 0; i < len(ratings); i++ {
		counts[i] = 1
	}

	// 從左到右，最左側往右比過去
	for i := 1; i < len(ratings); i++ {
		if ratings[i-1] < ratings[i] {
			counts[i] = counts[i-1] + 1
		}
	}

	// 從右到左，最右側往左比過去
	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i+1] < ratings[i] {
			counts[i] = max(counts[i], counts[i+1]+1)
		}
	}

	for i := 0; i < len(counts); i++ {
		res += counts[i]
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//leetcode submit region end(Prohibit modification and deletion)
