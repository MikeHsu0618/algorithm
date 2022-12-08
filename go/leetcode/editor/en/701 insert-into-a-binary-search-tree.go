package main

//You are given the root node of a binary search tree (BST) and a value to
//insert into the tree. Return the root node of the BST after the insertion. It is
//guaranteed that the new value does not exist in the original BST.
//
// Notice that there may exist multiple valid ways for the insertion, as long
//as the tree remains a BST after insertion. You can return any of them.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3], val = 5
//Output: [4,2,7,1,3,5]
//Explanation: Another accepted tree is:
//
//
//
// Example 2:
//
//
//Input: root = [40,20,60,10,30,50,70], val = 25
//Output: [40,20,60,10,30,50,70,null,null,25]
//
//
// Example 3:
//
//
//Input: root = [4,2,7,1,3,null,null,null,null,null,null], val = 5
//Output: [4,2,7,1,3,5]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree will be in the range [0, 10⁴].
// -10⁸ <= Node.val <= 10⁸
// All the values Node.val are unique.
// -10⁸ <= val <= 10⁸
// It's guaranteed that val does not exist in the original BST.
//
//
// Related Topics Tree Binary Search Tree Binary Tree 👍 4303 👎 154

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法三：優化解法一
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}
	if root.Val < val {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// 解法二：迭代 DFS 不分順序
func insertIntoBST2(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	cur := root
	var pre *TreeNode
	for {
		if cur == nil {
			if val > pre.Val {
				pre.Right = &TreeNode{Val: val}
			}
			if val < pre.Val {
				pre.Left = &TreeNode{Val: val}
			}
			break
		}
		pre = cur
		if cur.Val > val {
			cur = cur.Left
			continue
		}
		if cur.Val < val {
			cur = cur.Right
			continue
		}
	}
	return root
}

// 解法一： 遞迴 DFS 不分順序，遍歷到末端(node == nil)後直接可以將值插入適合的位置
// 利用 BTS 左邊的節點一定小於根節點與右邊節點一定大於根節點的特性，以 pre 前節點實現
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}
	var pre *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			if val > pre.Val {
				pre.Right = &TreeNode{Val: val}
			}
			if val < pre.Val {
				pre.Left = &TreeNode{Val: val}
			}
			return
		}
		pre = node
		if node.Val > val {
			travel(node.Left)
		}
		if node.Val < val {
			travel(node.Right)
		}
	}
	travel(root)
	return root
}

//leetcode submit region end(Prohibit modification and deletion)
