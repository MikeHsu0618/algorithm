package main

import "container/list"

//Given the root of a complete binary tree, return the number of the nodes in
//the tree.
//
// According to Wikipedia, every level, except possibly the last, is completely
//filled in a complete binary tree, and all nodes in the last level are as far
//left as possible. It can have between 1 and 2ʰ nodes inclusive at the last level h.
//
//
// Design an algorithm that runs in less than O(n) time complexity.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,4,5,6]
//Output: 6
//
//
// Example 2:
//
//
//Input: root = []
//Output: 0
//
//
// Example 3:
//
//
//Input: root = [1]
//Output: 1
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5 * 10⁴].
// 0 <= Node.val <= 5 * 10⁴
// The tree is guaranteed to be complete.
//
//
// Related Topics Binary Search Tree Depth-First Search Binary Tree 👍 6858 👎 3
//90

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法三：迭代 BFS
func countNodes3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			count++
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return count
}

// 解法二：遞迴版本二
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return getNodesCount(root)
}

func getNodesCount(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	count += getNodesCount(node.Left)
	count += getNodesCount(node.Right)
	return count + 1
}

// 解法一：遞迴不限遍歷順序
func countNodes1(root *TreeNode) int {
	count := 0
	var travelsal func(node *TreeNode)
	travelsal = func(node *TreeNode) {
		if node == nil {
			return
		}
		count++
		if node.Left != nil {
			travelsal(node.Left)
		}
		if node.Right != nil {
			travelsal(node.Right)
		}
	}
	travelsal(root)
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
