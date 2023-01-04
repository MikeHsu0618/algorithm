package main

//Given the root of a binary search tree and an integer k, return true if there
//exist two elements in the BST such that their sum is equal to k, or false
//otherwise.
//
//
// Example 1:
//
//
//Input: root = [5,3,6,2,4,null,7], k = 9
//Output: true
//
//
// Example 2:
//
//
//Input: root = [5,3,6,2,4,null,7], k = 28
//Output: false
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -10⁴ <= Node.val <= 10⁴
// root is guaranteed to be a valid binary search tree.
// -10⁵ <= k <= 10⁵
//
//
// Related Topics Hash Table Two Pointers Tree Depth-First Search Breadth-First
//Search Binary Search Tree Binary Tree 👍 5461 👎 236

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法一：看到 two sum 就要用 hash table，遍歷整顆樹時邊去查看 hash table 中有沒有另一個數值
func findTarget(root *TreeNode, k int) bool {
	m := make(map[int]bool, 0)
	var traversal func(node *TreeNode) bool
	traversal = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if m[k-node.Val] {
			return true
		}
		m[node.Val] = true
		return traversal(node.Left) || traversal(node.Right)
	}

	return traversal(root)
}

//leetcode submit region end(Prohibit modification and deletion)
