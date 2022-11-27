package main

import "container/list"

//Given a binary tree
//
//
//struct Node {
//  int val;
//  Node *left;
//  Node *right;
//  Node *next;
//}
//
//
// Populate each next pointer to point to its next right node. If there is no
//next right node, the next pointer should be set to NULL.
//
// Initially, all next pointers are set to NULL.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4,5,null,7]
//Output: [1,#,2,3,#,4,5,7,#]
//Explanation: Given the above binary tree (Figure A), your function should
//populate each next pointer to point to its next right node, just like in Figure B.
//The serialized output is in level order as connected by the next pointers, with
//'#' signifying the end of each level.
//
//
// Example 2:
//
//
//Input: root = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 6000].
// -100 <= Node.val <= 100
//
//
//
// Follow-up:
//
//
// You may only use constant extra space.
// The recursive approach is fine. You may assume implicit stack space does not
//count as extra space for this problem.
//
//
// Related Topics Linked List Tree Depth-First Search Breadth-First Search
//Binary Tree 👍 4978 👎 291

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// 與 116 解法一模一樣，差別在於是不是完全二叉樹
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			var node1, node2 *Node
			node1 = queue.Remove(queue.Front()).(*Node)
			if i != length-1 {
				node2 = queue.Front().Value.(*Node)
				node1.Next = node2
			} else {
				node1.Next = nil
			}
			if node1.Left != nil {
				queue.PushBack(node1.Left)
			}
			if node1.Right != nil {
				queue.PushBack(node1.Right)
			}
		}
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
