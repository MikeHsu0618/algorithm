package main

//You are given two binary trees root1 and root2.
//
// Imagine that when you put one of them to cover the other, some nodes of the
//two trees are overlapped while the others are not. You need to merge the two
//trees into a new binary tree. The merge rule is that if two nodes overlap, then sum
//node values up as the new value of the merged node. Otherwise, the NOT null
//node will be used as the node of the new tree.
//
// Return the merged tree.
//
// Note: The merging process must start from the root nodes of both trees.
//
//
// Example 1:
//
//
//Input: root1 = [1,3,2,5], root2 = [2,1,3,null,4,null,7]
//Output: [3,4,5,5,4,null,7]
//
//
// Example 2:
//
//
//Input: root1 = [1], root2 = [1,2]
//Output: [2,2]
//
//
//
// Constraints:
//
//
// The number of nodes in both trees is in the range [0, 2000].
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 76
//06 👎 266

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法二：優化解法一的判斷
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{Val: root1.Val + root2.Val}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

// 解法一：遞迴 DFS 前序，判斷 root1, root2 空值
func mergeTrees1(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	var val int
	var r1l, r1r, r2l, r2r *TreeNode
	if root1 == nil {
		r1l, r1r, r2l, r2r = nil, nil, root2.Left, root2.Right
		val = root2.Val
	} else if root2 == nil {
		r1l, r1r, r2l, r2r = root1.Left, root1.Right, nil, nil
		val = root1.Val
	} else {
		r1l, r1r, r2l, r2r = root1.Left, root1.Right, root2.Left, root2.Right
		val = root1.Val + root2.Val
	}
	root := &TreeNode{Val: val}
	root.Left = mergeTrees(r1l, r2l)
	root.Right = mergeTrees(r1r, r2r)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
