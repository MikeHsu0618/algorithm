package main

//Given the head of a linked list and an integer val, remove all the nodes of
//the linked list that has Node.val == val, and return the new head.
//
//
// Example 1:
//
//
//Input: head = [1,2,6,3,4,5,6], val = 6
//Output: [1,2,3,4,5]
//
//
// Example 2:
//
//
//Input: head = [], val = 1
//Output: []
//
//
// Example 3:
//
//
//Input: head = [7,7,7,7], val = 7
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the list is in the range [0, 10⁴].
// 1 <= Node.val <= 50
// 0 <= val <= 50
//
//
// Related Topics Linked List Recursion 👍 6305 👎 189

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用 dummyNode 觀念，並紀錄 head 指針位置(dummy.Next)，含後賦予額外變數遍歷比對
// 需要注意的是 return 必須是 dummy.Next 而不是 head, 才能處理 head 本身就是被刪除的情況
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	cur := dummy
	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
			continue
		}
		cur = cur.Next
	}
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
