package main

import "container/list"

//Given the root of a binary tree, return the inorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,3,2]
//
//
// Example 2:
//
//
//Input: root = []
//Output: []
//
//
// Example 3:
//
//
//Input: root = [1]
//Output: [1]
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
//
//Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics Stack Tree Depth-First Search Binary Tree 👍 10244 👎 485

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法三：使用統一迭代
// 中序遍歷：左中右
// 入棧順序：右中左
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back())
		if node == nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			continue
		}
		node1 := node.(*TreeNode)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		stack.PushBack(node1)
		stack.PushBack(nil)
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
	}
	return res
}

// 解法二：使用迭代（需要借助指針找到最底層）
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)

	stake := list.New()
	cur := root
	for cur != nil || stake.Len() > 0 {
		if cur != nil {
			stake.PushBack(cur)
			cur = cur.Left
			continue
		}
		// 到最底層之後開始往回迭代
		cur = stake.Remove(stake.Back()).(*TreeNode)
		res = append(res, cur.Val)
		cur = cur.Right
	}

	return res
}

// 解法一：使用遞迴(左中右)
func inorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	// 左中右
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
