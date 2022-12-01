package main

//Given a binary tree, determine if it is height-balanced.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: true
//
//
// Example 2:
//
//
//Input: root = [1,2,2,3,3,null,null,4,4]
//Output: false
//
//
// Example 3:
//
//
//Input: root = []
//Output: true
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 5000].
// -10⁴ <= Node.val <= 10⁴
//
//
// Related Topics Tree Depth-First Search Binary Tree 👍 7825 👎 435

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO 需要再回頭仔細了解
// 解法：使用遞迴檢查每一顆子樹是否平衡(左右高度最多不超過 1)
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeight := 1 + maxDepth(root.Left)   // 左
	rightHeight := 1 + maxDepth(root.Right) // 右
	// 中：檢查左右樹是否符合平衡樹
	return !(abs(leftHeight-rightHeight) > 1) && isBalanced(root.Left) && isBalanced(root.Right)
}

func maxDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(maxDepth(node.Left), maxDepth(node.Right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

//leetcode submit region end(Prohibit modification and deletion)
