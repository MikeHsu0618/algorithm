package main

//Given the root of a binary search tree (BST) with duplicates, return all the
//mode(s) (i.e., the most frequently occurred element) in it.
//
// If the tree has more than one mode, return them in any order.
//
// Assume a BST is defined as follows:
//
//
// The left subtree of a node contains only nodes with keys less than or equal
//to the node's key.
// The right subtree of a node contains only nodes with keys greater than or
//equal to the node's key.
// Both the left and right subtrees must also be binary search trees.
//
//
//
// Example 1:
//
//
//Input: root = [1,null,2,2]
//Output: [2]
//
//
// Example 2:
//
//
//Input: root = [0]
//Output: [0]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -10⁵ <= Node.val <= 10⁵
//
//
//
//Follow up: Could you do that without using any extra space? (Assume that the
//implicit stack space incurred due to recursion does not count).
//
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 2675
// 👎 621

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法二：遞迴 DFS 中序，使用 pre 技巧的搜尋樹特性
func findMode(root *TreeNode) []int {
	count := 0
	maxFreq := 0
	var pre *TreeNode
	res := make([]int, 0)
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if pre == nil {
			count = 1
		} else if pre.Val == node.Val {
			count++
		} else {
			count = 1
		}

		if count == maxFreq {
			res = append(res, node.Val)
		}

		if count > maxFreq {
			maxFreq = count
			res = append([]int{}, node.Val)
		}

		pre = node
		travel(node.Right)
	}
	travel(root)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
