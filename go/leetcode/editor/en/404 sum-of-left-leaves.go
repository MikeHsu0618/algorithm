package main

import "container/list"

//Given the root of a binary tree, return the sum of all left leaves.
//
// A leaf is a node with no children. A left leaf is a leaf that is the left
//child of another node.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: 24
//Explanation: There are two left leaves in the binary tree, with values 9 and 1
//5 respectively.
//
//
// Example 2:
//
//
//Input: root = [1]
//Output: 0
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 1000].
// -1000 <= Node.val <= 1000
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 40
//77 👎 265

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法三：迭代 DFS 前序
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count := 0
	stake := list.New()
	stake.PushBack(root)
	for stake.Len() > 0 {
		node := stake.Remove(stake.Back()).(*TreeNode)
		if node == nil {
			continue
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			count += node.Left.Val
		}
		stake.PushBack(node.Right)
		stake.PushBack(node.Left)
	}
	return count
}

// 解法二：優化解法一
func sumOfLeftLeaves2(root *TreeNode) int {
	return LeftNodesSum(root)
}

func LeftNodesSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
		count = node.Left.Val
	}
	return count + LeftNodesSum(node.Left) + LeftNodesSum(node.Right)
}

// 解法一：遞迴 DFS 前序
func sumOfLeftLeaves1(root *TreeNode) int {
	count := 0
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
			count += node.Left.Val
		}
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return count
}

//leetcode submit region end(Prohibit modification and deletion)
