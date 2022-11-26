package main

import "container/list"

//Given the root of a binary tree, return the bottom-up level order traversal
//of its nodes' values. (i.e., from left to right, level by level from leaf to root)
//.
//
//
// Example 1:
//
//
//Input: root = [3,9,20,null,null,15,7]
//Output: [[15,7],[9,20],[3]]
//
//
// Example 2:
//
//
//Input: root = [1]
//Output: [[1]]
//
//
// Example 3:
//
//
//Input: root = []
//Output: []
//
//
//
// Constraints:
//
//
// The number of nodes in the tree is in the range [0, 2000].
// -1000 <= Node.val <= 1000
//
//
// Related Topics Tree Breadth-First Search Binary Tree 👍 3921 👎 304

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
// 解法二：使用 queue 實現迭代 DFS 後反轉數組
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	depth := 0
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		// 紀錄 queue 長度代表該層的數量
		length := queue.Len()
		res = append(res, []int{})
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			res[depth] = append(res[depth], node.Val)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		depth++
	}
	reverse(res)
	return res
}

// 解法一：使用遞迴 DFS 後反轉數組
func levelOrderBottom1(root *TreeNode) [][]int {
	res := make([][]int, 0)
	depth := 0

	var order func(node *TreeNode, depth int)
	order = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}
		if len(res) == depth {
			res = append(res, []int{})
		}
		order(node.Left, depth+1)
		order(node.Right, depth+1)
		res[depth] = append(res[depth], node.Val)
	}
	order(root, depth)
	reverse(res)
	return res
}

func reverse(arr [][]int) {
	l, r := 0, len(arr)-1
	for l < r {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

//leetcode submit region end(Prohibit modification and deletion)
