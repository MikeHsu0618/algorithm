package main

//Given the root of a binary tree and an integer targetSum, return all root-to-
//leaf paths where the sum of the node values in the path equals targetSum. Each
//path should be returned as a list of the node values, not node references.
//
// A root-to-leaf path is a path starting from the root and ending at any leaf
//node. A leaf is a node with no children.
//
//
// Example 1:
//
//
//Input: root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//Output: [[5,4,11,2],[5,8,4,5]]
//Explanation: There are two paths whose sum equals targetSum:
//5 + 4 + 11 + 2 = 22
//5 + 8 + 4 + 5 = 22
//
//
// Example 2:
//
//
//Input: root = [1,2,3], targetSum = 5
//Output: []
//
//
// Example 3:
//
//
//Input: root = [1,2], targetSum = 0
//Output: []
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
// Related Topics Backtracking Tree Depth-First Search Binary Tree 👍 6496 👎 13
//3

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法：遞迴 DFS 前序（需注意 path []int 為指針，傳值時需特別處理）
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	var travel func(node *TreeNode, targetSum int, path []int)
	travel = func(node *TreeNode, targetSum int, path []int) {
		if node == nil {
			return
		}
		targetSum -= node.Val
		path = append(path, node.Val)
		if node.Left == nil && node.Right == nil && targetSum == 0 {
			res = append(res, append([]int{}, path...))
		}
		travel(node.Left, targetSum, path)
		travel(node.Right, targetSum, path)
	}
	travel(root, targetSum, []int{})
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
