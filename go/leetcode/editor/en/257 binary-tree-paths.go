package main

import "strconv"

//Given the root of a binary tree, return all root-to-leaf paths in any order.
//
// A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [1,2,3,null,5]
//Output: ["1->2->5","1->3"]
//
//
// Example 2:
//
//
//Input: root = [1]
//Output: ["1"]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics String Backtracking Tree Depth-First Search Binary Tree 👍 519
//9 👎 222

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO 回頭研究清楚
// 解法一：使用遞迴前序 DFS ＋ 回朔（隱藏在 s 當中）
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	var travel func(node *TreeNode, s string)
	travel = func(node *TreeNode, s string) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			res = append(res, s+strconv.Itoa(node.Val))
			return
		}
		s += strconv.Itoa(node.Val) + "->"
		travel(node.Left, s)
		travel(node.Right, s)
	}
	travel(root, "")
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
