package main

//Given a root node reference of a BST and a key, delete the node with the
//given key in the BST. Return the root node reference (possibly updated) of the BST.
//
// Basically, the deletion can be divided into two stages:
//
//
// Search for a node to remove.
// If the node is found, delete the node.
//
//
//
// Example 1:
//
//
//Input: root = [5,3,6,2,4,null,7], key = 3
//Output: [5,4,6,2,null,null,7]
//Explanation: Given key to delete is 3. So we find the node with value 3 and
//delete it.
//One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
//Please notice that another valid answer is [5,2,6,null,4,null,7] and it's
//also accepted.
//
//
//
// Example 2:
//
//
//Input: root = [5,3,6,2,4,null,7], key = 0
//Output: [5,3,6,2,4,null,7]
//Explanation: The tree does not contain a node with value = 0.
//
//
// Example 3:
//
//
//Input: root = [], key = 0
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 10⁴].
// -10⁵ <= Node.val <= 10⁵
// Each node has a unique value.
// root is a valid binary search tree.
// -10⁵ <= key <= 10⁵
//
//
//
// Follow up: Could you solve it with time complexity O(height of tree)?
//
// Related Topics Tree Binary Search Tree Binary Tree 👍 6673 👎 172

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO 要搞清楚
// https://www.delftstack.com/zh-tw/tutorial/data-structure/binary-search-tree-delete/
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 第一種情況：沒有找到刪除的點，遍歷到空節點直接返回
	if root == nil {
		return nil
	}

	if key < root.Val {
		root.Left = deleteNode(root.Left, key)
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
		return root
	}
	// 第二種情況：左右孩子都為空。直接刪除節點(可被第一種情況處理)

	// 第三種情況：左為空，右不為空，刪除節點，右孩子補位，返回根節點
	if root.Left == nil {
		return root.Right
	}
	// 第四種情況：右為空，左不為空，刪除節點，左節點補位，返回根節點
	if root.Right == nil {
		return root.Left
	}
	// 第五種情況：左右孩子都不為空
	minNode := findMinNode(root.Right)
	// 重新排序 1 : 將刪除節點的左孩子放到刪除節點的右子樹的最左面節點的左孩子位置
	minNode.Left = root.Left
	root = root.Right
	// 重新排序 2 : 將右子樹的最左面節點轉移到被刪除的節點上
	// root.Val = minNode.Val
	// root.Right = deleteNode(root.Right, minNode.Val)
	return root
}

func findMinNode(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

//leetcode submit region end(Prohibit modification and deletion)
