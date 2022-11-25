package main

import "container/list"

//Given the root of a binary tree, return the preorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [1,2,3]
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
// Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics Stack Tree Depth-First Search Binary Tree 👍 5388 👎 146

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

// 解法四：使用統一迭代的方法（再中間節點額外查入一個 nil 作為標示）
// 前序遍歷：中左右
// 入棧順序：右左中
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 彈出元素檢查是否為需要處理的中間節點(nil)
		node := stack.Remove(stack.Back())
		if node == nil {
			// 再次彈出，將中間節點放入結果集中
			node := stack.Remove(stack.Back()).(*TreeNode)
			res = append(res, node.Val)
			continue
		}
		// 入棧順序：右左中
		node1 := node.(*TreeNode)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
		stack.PushBack(node1)
		stack.PushBack(nil)
	}
	return res
}

// 解法三： 使用 linked-list 實現 stake 迭代
func preorderTraversal3(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	// 使用 root node 去初始化一個 doubly linked-list 的 stack
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		// 中左右 -> 中右左 (因為 stake 先進後出)
		res = append(res, node.Val)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return res
}

// 解法二：使用額外函數的遞迴
func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val) // middle
		traversal(node.Left)        // left
		traversal(node.Right)       // right
	}
	traversal(root)
	return res
}

// 解法一：使用沒有額外函式的遞迴
func preorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)            // middle
	left := preorderTraversal(root.Left)   // left
	right := preorderTraversal(root.Right) // right
	res = append(res, left...)
	res = append(res, right...)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
