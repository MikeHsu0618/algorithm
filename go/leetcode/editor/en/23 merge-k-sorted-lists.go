package main

//You are given an array of k linked-lists lists, each linked-list is sorted in
//ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Example 1:
//
//
//Input: lists = [[1,4,5],[1,3,4],[2,6]]
//Output: [1,1,2,3,4,4,5,6]
//Explanation: The linked-lists are:
//[
//  1->4->5,
//  1->3->4,
//  2->6
//]
//merging them into one sorted list:
//1->1->2->3->4->4->5->6
//
//
// Example 2:
//
//
//Input: lists = []
//Output: []
//
//
// Example 3:
//
//
//Input: lists = [[]]
//Output: []
//
//
//
// Constraints:
//
//
// k == lists.length
// 0 <= k <= 10⁴
// 0 <= lists[i].length <= 500
// -10⁴ <= lists[i][j] <= 10⁴
// lists[i] is sorted in ascending order.
// The sum of lists[i].length will not exceed 10⁴.
//
//
// Related Topics Linked List Divide and Conquer Heap (Priority Queue) Merge
//Sort 👍 14680 👎 550

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法一：不斷的將兩個 list 合併
// 優化 -> 使用分治法，將所有 list 兩兩一組合併，重複至只剩一組
func mergeKLists(lists []*ListNode) *ListNode {
	// 處理 edge case，lists 長度為零的情況
	if len(lists) == 0 {
		return nil
	}
	// 將第一個 list 作為起點，依序跟下一個合併
	result := lists[0]
	for i := 1; i < len(lists); i++ {
		result = mergeTwoLists(result, lists[i])
	}
	return result
}

// 解法二： 使用分治法優化解法一
func mergeKLists2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		// pop 2 lists
		l1 := lists[0]
		l2 := lists[1]
		lists = lists[2:]

		merged := mergeTwoLists(l1, l2)
		lists = append(lists, merged)
	}

	return lists[0]
}

//leetcode submit region end(Prohibit modification and deletion)
