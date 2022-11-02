package main

//You are given the heads of two sorted linked lists list1 and list2.
//
// Merge the two lists in a one sorted list. The list should be made by
//splicing together the nodes of the first two lists.
//
// Return the head of the merged linked list.
//
//
// Example 1:
//
//
//Input: list1 = [1,2,4], list2 = [1,3,4]
//Output: [1,1,2,3,4,4]
//
//
// Example 2:
//
//
//Input: list1 = [], list2 = []
//Output: []
//
//
// Example 3:
//
//
//Input: list1 = [], list2 = [0]
//Output: [0]
//
//
//
// Constraints:
//
//
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.
//
//
// Related Topics Linked List Recursion 👍 15703 👎 1380

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 遞歸操作當某一個list為空的時候，退出
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// 如果其中一個 list 為 nil 則後續為另一個 list
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	// 比較兩個 list 當前值大小來決定合併順序，並且繼續遞迴比較
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}

//leetcode submit region end(Prohibit modification and deletion)
