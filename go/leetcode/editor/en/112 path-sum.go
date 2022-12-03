package main

import "container/list"

//Given the root of a binary tree and an integer targetSum, return true if the
//tree has a root-to-leaf path such that adding up all the values along the path
//equals targetSum.
//
// A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,null,1], targetSum = 22
//Output: true
//Explanation: The root-to-leaf path with the target sum is shown.
//
//
// Example 2:
//
//
//Input: root = [1,2,3], targetSum = 5
//Output: false
//Explanation: There two root-to-leaf paths in the tree:
//(1 --> 2): The sum is 3.
//(1 --> 3): The sum is 4.
//There is no root-to-leaf path with sum = 5.
//
//
// Example 3:
//
//
//Input: root = [], targetSum = 0
//Output: false
//Explanation: Since the tree is empty, there are no root-to-leaf paths.
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5000].
// -1000 <= Node.val <= 1000
// -1000 <= targetSum <= 1000
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 75
//24 👎 901

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type Item struct {
	node *TreeNode
	val  int
}

// 解法二：隱藏回溯的迭代 DFS
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := list.New()
	stack.PushBack(Item{root, root.Val})
	for stack.Len() > 0 {
		item := stack.Remove(stack.Back()).(Item)
		if item.node.Left == nil && item.node.Right == nil && item.val == targetSum {
			return true
		}
		if item.node.Right != nil {
			stack.PushBack(Item{item.node.Right, item.val + item.node.Right.Val})
		}
		if item.node.Left != nil {
			stack.PushBack(Item{item.node.Left, item.val + item.node.Left.Val})
		}
	}
	return false
}

// 解法一：隱藏回朔的遞迴 DFS
func hasPathSum1(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil && targetSum == 0 {
		return true
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}

//leetcode submit region end(Prohibit modification and deletion)
