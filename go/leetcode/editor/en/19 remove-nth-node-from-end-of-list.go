package main

//Given the head of a linked list, remove the nᵗʰ node from the end of the list
//and return its head.
//
//
// Example 1:
//
//
//Input: head = [1,2,3,4,5], n = 2
//Output: [1,2,3,5]
//
//
// Example 2:
//
//
//Input: head = [1], n = 1
//Output: []
//
//
// Example 3:
//
//
//Input: head = [1,2], n = 1
//Output: [1]
//
//
//
// Constraints:
//
//
// The number of nodes in the list is sz.
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
//
//
//
// Follow up: Could you do this in one pass?
//
// Related Topics Linked List Two Pointers 👍 13624 👎 562

// ListNode leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 解法三：優化解法二一點點
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for fast.Next != nil {
		fast = fast.Next
		if n > 0 {
			n--
			continue
		}
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 解法二： 使用快慢指針＋ dummy node 的概念，首先將快指針從首部移動到第 n 個距離
// 使雙指針同時移動，直到快指針觸碰到盡頭，則慢指針的下一位數值即為倒數第 n 個
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy
	for count := 1; fast.Next != nil; count++ {
		fast = fast.Next
		if count > n {
			slow = slow.Next
		}
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

//leetcode submit region end(Prohibit modification and deletion)
// 解法一: 簡單遍歷出長度，並再次遍歷剃除倒數第 n 個
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	// 先遍歷一次 node 來取得總長度
	length := 0
	for tmp := head; tmp != nil; {
		length++
		tmp = tmp.Next
	}
	// 處理 eage case n === 1 的情況
	if n == length {
		return head.Next
	}

	// 遍歷到要剃除的前一個時，略過下一位數值
	list := head
	for count := 1; list != nil; count++ {
		if count == length-n {
			list.Next = list.Next.Next
		}
		list = list.Next
	}
	return head
}
