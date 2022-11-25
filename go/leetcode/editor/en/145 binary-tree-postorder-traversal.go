package main

import "container/list"

//Given the root of a binary tree, return the postorder traversal of its nodes'
//values.
//
//
// Example 1:
//
//
//Input: root = [1,null,2,3]
//Output: [3,2,1]
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
// The number of the nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
//
//Follow up: Recursive solution is trivial, could you do it iteratively?
//
// Related Topics Stack Tree Depth-First Search Binary Tree 👍 5263 👎 159

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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法三：使用統一迭代
// 後序遍歷：左右中
// 入棧順序：中右左
func postorderTraversal(root *TreeNode) []int {
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
		stack.PushBack(node1)
		stack.PushBack(nil)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
	}

	return res
}

// 解法二：使用前序迭代並且反轉數列，即可得到後序遍歷
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		// 中左右(前序) -> 中右左(經 stack 先進後出後) -> 左右中（反轉後）
		node := stack.Remove(stack.Back()).(*TreeNode)
		res = append(res, node.Val)
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
	}
	reverse(res)
	return res
}

func reverse(nums []int) {
	l, r := 0, len(nums)-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// 解法一：使用遞迴
func postorderTraversal1(root *TreeNode) []int {
	res := make([]int, 0)
	var traversal func(node *TreeNode)
	// 左右中
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)
	}
	traversal(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
