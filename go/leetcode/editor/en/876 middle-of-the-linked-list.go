package main

//Given the head of a singly linked list, return the middle node of the linked
//list.
//
// If there are two middle nodes, return the second middle node.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5]
//Output: [3,4,5]
//Explanation: The middle node of the list is node 3.
//
//
// Example 2:
//
//
//Input: head = [1,2,3,4,5,6]
//Output: [4,5,6]
//Explanation: Since the list has two middle nodes with values 3 and 4, we
//return the second one.
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [1, 100].
// 1 <= Node.val <= 100
//
//
// Related Topics Linked List Two Pointers 👍 8379 👎 234

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 快指針走兩部
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 先算出長度再找出中間點
func middleNode1(head *ListNode) *ListNode {
	length := 0
	cur := head

	for cur != nil {
		cur = cur.Next
		length++
	}
	cur = head
	for i := 0; i < length/2; i++ {
		cur = cur.Next
	}

	return cur
}

//leetcode submit region end(Prohibit modification and deletion)
