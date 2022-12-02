package main

import "container/list"

//Given the root of a binary tree, return the leftmost value in the last row of
//the tree.
//
//
// Example 1:
//
//
//Input: root = [2,1,3]
//Output: 1
//
//
// Example 2:
//
//
//Input: root = [1,2,3,4,null,5,6,null,null,7]
//Output: 7
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 26
//40 👎 235

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法二：遞迴 DFS 前序＋回朔
func findBottomLeftValue(root *TreeNode) int {
	leftest := -1
	maxDepth := -1
	var traversal func(node *TreeNode, depth int)
	traversal = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil && depth > maxDepth {
			leftest = node.Val
			maxDepth = depth
		}
		traversal(node.Left, depth+1)
		traversal(node.Right, depth+1)
	}
	traversal(root, 0)
	return leftest
}

// 解法一：迭代 BFS 紀錄每層最左邊的數值
func findBottomLeftValue1(root *TreeNode) int {
	leftest := -1
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if i == 0 {
				leftest = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
	}
	return leftest
}

//leetcode submit region end(Prohibit modification and deletion)
