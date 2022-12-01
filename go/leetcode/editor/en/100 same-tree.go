package main

//Given the roots of two binary trees p and q, write a function to check if
//they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical,
//and the nodes have the same value.
//
//
// Example 1:
//
//
//Input: p = [1,2,3], q = [1,2,3]
//Output: true
//
//
// Example 2:
//
//
//Input: p = [1,2], q = [1,null,2]
//Output: false
//
//
// Example 3:
//
//
//Input: p = [1,2,1], q = [1,1,2]
//Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in both trees is in the range [0, 100].
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 76
//82 👎 162

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//  解法三：優化解法二
func isSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// 解法二：遞迴 DFS 前序
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	var compare func(left, right *TreeNode) bool
	compare = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return compare(left.Left, right.Left) && compare(left.Right, right.Right)
	}
	return compare(p, q)
}

// 解法一：迭代 DFS 前序
func isSameTree1(p *TreeNode, q *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, p, q)
	for len(queue) > 0 {
		left, right := queue[0], queue[1]
		queue = queue[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		queue = append(queue, left.Left, right.Left, left.Right, right.Right)
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
