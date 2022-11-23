package main

import "container/heap"

//Given an integer array nums and an integer k, return the k most frequent
//elements. You may return the answer in any order.
//
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
//Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
//Output: [1]
//
//
// Constraints:
//
//
// 1 <= nums.length <= 10⁵
// -10⁴ <= nums[i] <= 10⁴
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
//
//
//
// Follow up: Your algorithm's time complexity must be better than O(n log n),
//where n is the array's size.
//
// Related Topics Array Hash Table Divide and Conquer Sorting Heap (Priority
//Queue) Bucket Sort Counting Quickselect 👍 11909 👎 438

//leetcode submit region begin(Prohibit modification and deletion)
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	pq := make(PriorityQueue, 0)
	for val, amount := range m {
		pq = append(pq, &Item{
			Val:    val,
			Amount: amount,
		})
	}
	heap.Init(&pq)
	for k > 0 {
		item := heap.Pop(&pq)
		res = append(res, item.(*Item).Val)
		k--
	}
	return res
}

type Item struct {
	Val    int
	Amount int
}

type PriorityQueue []*Item

func (pq *PriorityQueue) Len() int {
	return len(*pq)
}

// 實現最大堆
func (pq *PriorityQueue) Less(i, j int) bool {
	return (*pq)[i].Amount > (*pq)[j].Amount
}

func (pq *PriorityQueue) Swap(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*Item))
}

func (pq *PriorityQueue) Pop() interface{} {
	res := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
