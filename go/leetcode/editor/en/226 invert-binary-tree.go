package main

import "container/list"

//Given the root of a binary tree, invert the tree, and return its root.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3,6,9]
//Output: [4,7,2,9,6,3,1]
//
//
// Example 2:
//
//
//Input: root = [2,1,3]
//Output: [2,3,1]
//
//
// Example 3:
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
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 10
//406 👎 146

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
		}
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
