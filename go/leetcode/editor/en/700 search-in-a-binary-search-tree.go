package main

import "container/list"

//You are given the root of a binary search tree (BST) and an integer val.
//
// Find the node in the BST that the node's value equals val and return the
//subtree rooted with that node. If such a node does not exist, return null.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3], val = 2
//Output: [2,1,3]
//
//
// Example 2:
//
//
//Input: root = [4,2,7,1,3], val = 5
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 5000].
// 1 <= Node.val <= 10⁷
// root is a binary search tree.
// 1 <= val <= 10⁷
//
//
// Related Topics Tree Binary Search Tree Binary Tree 👍 4328 👎 159

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法三：迭代 BFS 利用特性
func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val > val {
			root = root.Left
			continue
		}
		if root.Val < val {
			root = root.Right
			continue
		}
		return root
	}
	return root
}

// 解法二：遞迴 DFS 利用特性
func searchBST2(root *TreeNode, val int) *TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if root.Val > val {
		return searchBST(root.Left, val)
	}
	if root.Val < val {
		return searchBST(root.Right, val)
	}
	return nil
}

// 解法一：迭代 BFS 沒有利用到 binary search tree 的特性
func searchBST1(root *TreeNode, val int) *TreeNode {
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node == nil {
				continue
			}
			if node.Val == val {
				return node
			}
			queue.PushBack(node.Left)
			queue.PushBack(node.Right)
		}
	}
	return nil
}

//leetcode submit region end(Prohibit modification and deletion)
