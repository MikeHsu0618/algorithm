package main

//Given head, the head of a linked list, determine if the linked list has a
//cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can
//be reached again by continuously following the next pointer. Internally, pos is
//used to denote the index of the node that tail's next pointer is connected to.
//Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
//
//
// Example 1:
//
//
//Input: head = [3,2,0,-4], pos = 1
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to
//the 1st node (0-indexed).
//
//
// Example 2:
//
//
//Input: head = [1,2], pos = 0
//Output: true
//Explanation: There is a cycle in the linked list, where the tail connects to
//the 0th node.
//
//
// Example 3:
//
//
//Input: head = [1], pos = -1
//Output: false
//Explanation: There is no cycle in the linked list.
//
//
//
// Constraints:
//
//
// The number of the nodes in the list is in the range [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// pos is -1 or a valid index in the linked-list.
//
//
//
// Follow up: Can you solve it using O(1) (i.e. constant) memory?
//
// Related Topics Hash Table Linked List Two Pointers 👍 10971 👎 930

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 使用快慢指針，慢指針每次走一步快指針每次走兩步，遲早會遇到
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	// 需要檢查零界值
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}

// hash table O(n) O(n)
func hasCycle1(head *ListNode) bool {
	m := make(map[*ListNode]int, 0)

	cur := head
	for cur != nil {
		if _, ok := m[cur]; ok {
			return true
		}
		m[cur]++
		cur = cur.Next
	}

	return false
}

//leetcode submit region end(Prohibit modification and deletion)
