package main

//Given a linked list, swap every two adjacent nodes and return its head. You
//must solve the problem without modifying the values in the list's nodes (i.e.,
//only nodes themselves may be changed.)
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4]
//Output: [2,1,4,3]
//
//
// Example 2:
//
//
//Input: head = []
//Output: []
//
//
// Example 3:
//
//
//Input: head = [1]
//Output: [1]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 100].
// 0 <= Node.val <= 100
//
//
// Related Topics Linked List Recursion 👍 8477 👎 337

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// dummy -> 1 -> 2 -> 3
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur.Next != nil && cur.Next.Next != nil {
		next := cur.Next.Next.Next
		first := cur.Next
		// dummy -> 2
		cur.Next = cur.Next.Next
		// 2 -> 1
		cur.Next.Next = first
		// 1 -> 3
		cur.Next.Next.Next = next
		cur = cur.Next.Next
	}
	return dummy.Next
}

// dummy -> 1 -> 2 -> 3 => dummy -> 2 -> 1 -> 3
func swapPairs1(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first, second, third := cur.Next, cur.Next.Next, cur.Next.Next.Next
		cur.Next = second
		cur.Next.Next = first
		cur.Next.Next.Next = third
		cur = cur.Next.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
