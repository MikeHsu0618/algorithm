package main

//Given a binary search tree (BST), find the lowest common ancestor (LCA) node
//of two given nodes in the BST.
//
// According to the definition of LCA on Wikipedia: “The lowest common ancestor
//is defined between two nodes p and q as the lowest node in T that has both p
//and q as descendants (where we allow a node to be a descendant of itself).”
//
//
// Example 1:
//
//
//Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
//Output: 6
//Explanation: The LCA of nodes 2 and 8 is 6.
//
//
// Example 2:
//
//
//Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
//Output: 2
//Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant
//of itself according to the LCA definition.
//
//
// Example 3:
//
//
//Input: root = [2,1], p = 2, q = 1
//Output: 2
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [2, 10⁵].
// -10⁹ <= Node.val <= 10⁹
// All Node.val are unique.
// p != q
// p and q will exist in the BST.
//
//
// Related Topics Tree Depth-First Search Binary Search Tree Binary Tree 👍 8413
// 👎 241

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法二：迭代 DFS 後序
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for {
		if root.Val < p.Val && root.Val < q.Val {
			root = root.Right
			continue
		}
		if root.Val > p.Val && root.Val > q.Val {
			root = root.Left
			continue
		}
		return root
	}
	return root
}

// 解法一：遞迴 DFS 後序，使用 BST 特性，由下往上找，找到的第一個即返回
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
