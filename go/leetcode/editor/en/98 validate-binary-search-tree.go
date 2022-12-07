package main

import "math"

//Given the root of a binary tree, determine if it is a valid binary search
//tree (BST).
//
// A valid BST is defined as follows:
//
//
// The left subtree of a node contains only nodes with keys less than the
//node's key.
// The right subtree of a node contains only nodes with keys greater than the
//node's key.
// Both the left and right subtrees must also be binary search trees.
//
//
//
// Example 1:
//
//
//Input: root = [2,1,3]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [5,1,4,null,null,3,6]
//Output: false
//Explanation: The root node's value is 5 but its right child's value is 4.
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
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 1328
//9 👎 1073

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法二：遞迴 DFS 中序 (左右中的遍歷順序剛好就符合二元搜尋樹)
func isValidBST(root *TreeNode) bool {
	var pre *TreeNode
	var isValid func(node *TreeNode) bool
	isValid = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		left := isValid(node.Left)
		if pre != nil && node.Val <= pre.Val {
			return false
		}
		pre = node
		right := isValid(node.Right)
		return left && right
	}
	return isValid(root)
}

// 解法一：遞迴
func isValidBST(root *TreeNode) bool {
	return check(root, math.MinInt64, math.MaxInt64)
}

func check(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	if min >= int64(node.Val) || max <= int64(node.Val) {
		return false
	}
	return check(node.Left, min, int64(node.Val)) && check(node.Right, int64(node.Val), max)
}

//leetcode submit region end(Prohibit modification and deletion)
