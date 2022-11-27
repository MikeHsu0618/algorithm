package main

import (
	"container/list"
	"math"
)

//Given the root of a binary tree, return an array of the largest value in each
//row of the tree (0-indexed).
//
//
// Example 1:
//
//
//Input: root = [1,3,2,5,3,null,9]
//Output: [1,3,9]
//
//
// Example 2:
//
//
//Input: root = [1,2,3]
//Output: [1,3]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree will be in the range [0, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 24
//32 👎 92

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func largestValues(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		max := math.MinInt64
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Val > max {
				max = node.Val
			}
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, int(max))
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
