package main

//Given the head of a linked list, return the node where the cycle begins. If
//there is no cycle, return null.
//
// There is a cycle in a linked list if there is some node in the list that can
//be reached again by continuously following the next pointer. Internally, pos is
//used to denote the index of the node that tail's next pointer is connected to (0
//-indexed). It is -1 if there is no cycle. Note that pos is not passed as a
//parameter.
//
// Do not modify the linked list.
//
//
// Example 1:
//
//
//Input: head = [3,2,0,-4], pos = 1
//Output: tail connects to node index 1
//Explanation: There is a cycle in the linked list, where tail connects to the
//second node.
//
//
// Example 2:
//
//
//Input: head = [1,2], pos = 0
//Output: tail connects to node index 0
//Explanation: There is a cycle in the linked list, where tail connects to the
//first node.
//
//
// Example 3:
//
//
//Input: head = [1], pos = -1
//Output: no cycle
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
// Related Topics Hash Table Linked List Two Pointers 👍 9397 👎 661

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 解法二(雙指針)：使用快慢指針，fast走兩步，slow走一步直到相會，head 與將會點開始同時接近直到再次會合
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for slow != head {
				slow = slow.Next
				head = head.Next
			}
			return slow
		}
	}
	return nil
}

// 解法一(暴力解)： 遍歷整個 list 並存進 map 中，直到遇到重複為止
func detectCycle1(head *ListNode) *ListNode {
	m := make(map[*ListNode]int)
	for head != nil {
		if _, ok := m[head]; ok {
			break
		}
		m[head]++
		head = head.Next
	}
	return head
}

//leetcode submit region end(Prohibit modification and deletion)
