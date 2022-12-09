package main

//Given two integers n and k, return all possible combinations of k numbers
//chosen from the range [1, n].
//
// You may return the answer in any order.
//
//
// Example 1:
//
//
//Input: n = 4, k = 2
//Output: [[1,2],[1,3],[1,4],[2,3],[2,4],[3,4]]
//Explanation: There are 4 choose 2 = 6 total combinations.
//Note that combinations are unordered, i.e., [1,2] and [2,1] are considered to
//be the same combination.
//
//
// Example 2:
//
//
//Input: n = 1, k = 1
//Output: [[1]]
//Explanation: There is 1 choose 1 = 1 total combination.
//
//
//
// Constraints:
//
//
// 1 <= n <= 20
// 1 <= k <= n
//
//
// Related Topics Backtracking 👍 5400 👎 170

//leetcode submit region begin(Prohibit modification and deletion)
// for (int i = startIndex; i <= n; i++) {
// 接下来看一下优化过程如下：
// 已经选择的元素个数：path.size();
// 所需需要的元素个数为: k - path.size();
// 列表中剩余元素（n-i） >= 所需需要的元素个数（k - path.size()）
// 在集合n中至多要从该起始位置 : i <= n - (k - path.size()) + 1，开始遍历
// 为什么有个+1呢，因为包括起始位置，我们要是一个左闭的集合。
// 举个例子，n = 4，k = 3， 目前已经选取的元素为0（path.size为0），n - (k - 0) + 1 即 4 - ( 3 - 0) + 1 = 2。
// 从2开始搜索都是合理的，可以是组合[2, 3, 4]。
// for (int i = startIndex; i <= n - (k - path.size()) + 1; i++) // i为本次搜索的起始位置
func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)
	var backTracing func(start int)
	backTracing = func(start int) {
		if len(path) == k {
			res = append(res, append([]int{}, path...))
			return
		}
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			backTracing(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTracing(1)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
