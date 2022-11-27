package main

import "container/list"

//Given the root of a binary tree, invert the tree, and return its root.
//
//
// Example 1:
//
//
//Input: root = [4,2,7,1,3,6,9]
//Output: [4,7,2,9,6,3,1]
//
//
// Example 2:
//
//
//Input: root = [2,1,3]
//Output: [2,3,1]
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
// The number of nodes in the tree is in the range [0, 100].
// -100 <= Node.val <= 100
//
//
// Related Topics Tree Depth-First Search Breadth-First Search Binary Tree 👍 10
//406 👎 146

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法八：統一 迭代 中序(左中右) DFS
// 入棧順序：右中左
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back())
		if node == nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			continue
		}
		node1 := node.(*TreeNode)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		stack.PushBack(node1)
		stack.PushBack(nil)
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
	}

	return root
}

// 解法七：統一 迭代 後序(左右中) DFS
// 入棧順序：中右左
func invertTree7(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back())
		if node == nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			continue
		}
		node1 := node.(*TreeNode)
		stack.PushBack(node1)
		stack.PushBack(nil)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
	}
	return root
}

// 解法六：統一 迭代 前序(中左右) DFS
// 入棧順序：右左中
func invertTree6(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back())
		if node == nil {
			node := stack.Remove(stack.Back()).(*TreeNode)
			node.Left, node.Right = node.Right, node.Left
			continue
		}
		node1 := node.(*TreeNode)
		if node1.Right != nil {
			stack.PushBack(node1.Right)
		}
		if node1.Left != nil {
			stack.PushBack(node1.Left)
		}
		stack.PushBack(node1)
		stack.PushBack(nil)
	}
	return root
}

// 解法五：迭代 後序(左右中) DFS
// 入棧順序：中右左
func invertTree5(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		node.Left, node.Right = node.Right, node.Left
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
	}
	return root
}

// 解法四：迭代 前序(中左右) DFS
// 入棧順序：右左中
func invertTree4(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	stack := list.New()
	stack.PushBack(root)
	for stack.Len() > 0 {
		node := stack.Remove(stack.Back()).(*TreeNode)
		if node.Right != nil {
			stack.PushBack(node.Right)
		}
		if node.Left != nil {
			stack.PushBack(node.Left)
		}
		node.Left, node.Right = node.Right, node.Left
	}
	return root
}

// 解法三：遞迴 後序(左右中) DFS
func invertTree3(root *TreeNode) *TreeNode {
	var travelsal func(root *TreeNode)
	travelsal = func(node *TreeNode) {
		if node == nil {
			return
		}
		travelsal(node.Left)
		travelsal(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	travelsal(root)
	return root
}

// 解法二：遞迴 前序(中左右) DFS
func invertTree2(root *TreeNode) *TreeNode {
	var travelsal func(root *TreeNode)
	travelsal = func(node *TreeNode) {
		if node == nil {
			return
		}
		node.Left, node.Right = node.Right, node.Left
		travelsal(node.Left)
		travelsal(node.Right)
	}
	travelsal(root)
	return root
}

// 解法一：迭代 BFS
func invertTree1(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			node.Left, node.Right = node.Right, node.Left
		}
	}

	return root
}

//leetcode submit region end(Prohibit modification and deletion)
