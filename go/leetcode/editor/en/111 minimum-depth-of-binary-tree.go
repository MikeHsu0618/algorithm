package main

import "container/list"

//Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the
//root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 2
//
//
// Example 2:
//
//
//Input: root = [2,null,3,null,4,null,5,null,6]
//Output: 5
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁵].
// -1000 <= Node.val <= 1000
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 51
//04 👎 1043

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法二：遞迴
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left != nil && root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	if root.Right != nil && root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	return 1 + min(minDepth(root.Right), minDepth(root.Left))
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 解法一：使用迭代 BFS，而 node.Left node.Right 都為空的值就是最小深度
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		length := queue.Len()
		depth++
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left == nil && node.Right == nil {
				return depth
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
