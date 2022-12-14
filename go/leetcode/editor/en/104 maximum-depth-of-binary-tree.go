package main

import "container/list"

//Given the root of a binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path
//from the root node down to the farthest leaf node.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 3
//
//
// Example 2:
//
//
//Input: root = [1,null,2]
//Output: 2
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁴].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 90
//66 👎 149

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 重點：前序可以算出從上往下算出深度，後序可以由下往上算出高度

// 解法三：遞迴 DFS 前序
func maxDepth3(root *TreeNode) int {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftDepth := 1 + getDepth(node.Left)
		rightDepth := 1 + getDepth(node.Right)
		return max(leftDepth, rightDepth)
	}
	return getDepth(root)
}

// 解法二：遞迴 DFS 後序
func maxDepth2(root *TreeNode) int {
	return getDepth(root)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	leftDepth := getDepth(node.Left)
	rightDepth := getDepth(node.Right)
	depth := 1 + max(leftDepth, rightDepth)
	return depth
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 解法一：迭代 BFS 簡單算出高度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 0
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
		}
		depth++
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
