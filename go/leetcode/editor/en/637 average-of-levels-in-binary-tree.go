package main

import "container/list"

//Given the root of a binary tree, return the average value of the nodes on
//each level in the form of an array. Answers within 10⁻⁵ of the actual answer will
//be accepted.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: [3.00000,14.50000,11.00000]
//Explanation: The average value of nodes on level 0 is 3, on level 1 is 14.5,
//and on level 2 is 11.
//Hence return [3, 14.5, 11].
//
//
// Example 2:
//
//
//Input: root = [3,9,20,15,7]
//Output: [3.00000,14.50000,11.00000]
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [1, 10⁴].
// -2³¹ <= Node.val <= 2³¹ - 1
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 44
//15 👎 276

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 解法：使用迭代 BFS
func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return nil
	}
	res := make([]float64, 0)
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		sum := 0
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			sum += node.Val
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		res = append(res, float64(sum)/float64(length))
	}
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
