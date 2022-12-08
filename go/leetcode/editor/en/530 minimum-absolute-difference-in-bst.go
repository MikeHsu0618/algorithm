package main

import (
	"fmt"
	"math"
)

//Given the root of a Binary Search Tree (BST), return the minimum absolute
//difference between the values of any two different nodes in the tree.
//
//
// Example 1:
//
//
//Input: root = [4,2,6,1,3]
//Output: 1
//
//
// Example 2:
//
//
//Input: root = [1,0,48,null,null,12,49]
//Output: 1
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [2, 10⁴].
// 0 <= Node.val <= 10⁵
//
//
//
// Note: This question is the same as 783: https://leetcode.com/problems/
//minimum-distance-between-bst-nodes/
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Search
//Tree Binary Tree 👍 2518 👎 135

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法一：遞迴 DFS 中序，保留前個節點並利用 BST 特性 ＋ 中序特性，確保遍歷節點數值由小到大
func getMinimumDifference1(root *TreeNode) int {
	minAb := math.MaxInt
	var pre *TreeNode
	var travel func(node *TreeNode)
	travel = func(node *TreeNode) {
		if node == nil {
			return
		}
		travel(node.Left)
		if pre != nil && node.Val-pre.Val < minAb {
			fmt.Println(pre.Val, node.Val)
			minAb = node.Val - pre.Val
		}
		pre = node
		travel(node.Right)
	}
	travel(root)
	return minAb
}

//leetcode submit region end(Prohibit modification and deletion)
