package main

//You are given an array of integers nums, there is a sliding window of size k
//which is moving from the very left of the array to the very right. You can only
//see the k numbers in the window. Each time the sliding window moves right by one
//position.
//
// Return the max sliding window.
//
//
// Example 1:
//
//
//Input: nums = [1,3,-1,-3,5,3,6,7], k = 3
//Output: [3,3,5,5,6,7]
//Explanation:
//Window position                Max
//---------------               -----
//[1  3  -1] -3  5  3  6  7       3
// 1 [3  -1  -3] 5  3  6  7       3
// 1  3 [-1  -3  5] 3  6  7       5
// 1  3  -1 [-3  5  3] 6  7       5
// 1  3  -1  -3 [5  3  6] 7       6
// 1  3  -1  -3  5 [3  6  7]      7
//
//
// Example 2:
//
//
//Input: nums = [1], k = 1
//Output: [1]
//
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// 1 <= k <= nums.length
//
//
// Related Topics Array Queue Sliding Window Heap (Priority Queue) Monotonic
//Queue 👍 12856 👎 418

//leetcode submit region begin(Prohibit modification and deletion)
// 有序隊列(單調數列)：維持遞增或遞減的隊列
type MyQueue struct {
	queue []int
}

func NewMyQueue() *MyQueue {
	return &MyQueue{}
}

func (q *MyQueue) Pop(val int) {
	// 關鍵：如果值不再滑窗範圍內著排除
	if !q.Empty() && val == q.Front() {
		q.queue = q.queue[1:len(q.queue)]
	}
}

func (q *MyQueue) Push(val int) {
	// 關鍵：只將需要用到的值保留在隊列中
	for !q.Empty() && val > q.Back() {
		q.queue = q.queue[:len(q.queue)-1]
	}
	q.queue = append(q.queue, val)
}

func (q *MyQueue) Front() int {
	return q.queue[0]
}

func (q *MyQueue) Back() int {
	return q.queue[len(q.queue)-1]
}

func (q *MyQueue) Empty() bool {
	return len(q.queue) == 0
}

func maxSlidingWindow(nums []int, k int) []int {
	res := make([]int, 0)
	que := NewMyQueue()
	for i := 0; i < k; i++ {
		que.Push(nums[i])
	}
	res = append(res, que.Front())
	// 開始滑窗
	for i := k; i < len(nums); i++ {
		que.Pop(nums[i-k])
		que.Push(nums[i])
		res = append(res, que.Front())
	}
	return res
}

// 暴力解：Time Limit Exceeded
func maxSlidingWindow1(nums []int, k int) []int {
	res := make([]int, 0)
	for i := 0; i < len(nums)-k+1; i++ {
		max := -10000
		for j := 0; j < k; j++ {
			if nums[i+j] > max {
				max = nums[i+j]
			}
		}
		res = append(res, max)
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
