package main

//Given the head of a singly linked list, reverse the list, and return the
//reversed list.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5]
//Output: [5,4,3,2,1]
//
//
// Example 2:
//
//
//Input: head = [1,2]
//Output: [2,1]
//
//
// Example 3:
//
//
//Input: head = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is the range [0, 5000].
// -5000 <= Node.val <= 5000
//
//
//
// Follow up: A linked list can be reversed either iteratively or recursively.
//Could you implement both?
//
// Related Topics Linked List Recursion 👍 15417 👎 258

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法二：使用遞迴優化解法一
func reverseList(head *ListNode) *ListNode {
	return reverse(nil, head)
}

func reverse(pre, head *ListNode) *ListNode {
	if head == nil {
		return pre
	}
	next := head.Next
	head.Next = pre
	return reverse(head, next)
}

// 解法一：將每個 node 的 Next 轉移到上一個 node 即代表完成反轉
func reverseList1(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

//leetcode submit region end(Prohibit modification and deletion)
