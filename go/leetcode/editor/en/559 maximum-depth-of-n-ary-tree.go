package main

import "container/list"

//Given a n-ary tree, find its maximum depth.
//
// The maximum depth is the number of nodes along the longest path from the
//root node down to the farthest leaf node.
//
// Nary-Tree input serialization is represented in their level order traversal,
//each group of children is separated by the null value (See examples).
//
//
// Example 1:
//
//
//
//
//Input: root = [1,null,3,2,4,null,5,6]
//Output: 3
//
//
// Example 2:
//
//
//
//
//Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,
//null,12,null,13,null,null,14]
//Output: 5
//
//
//
// Constraints:
//
//
// The total number of nodes is in the range [0, 10⁴].
// The depth of the n-ary tree is less than or equal to 1000.
//
//
// Related Topics Tree Depth-First Search Breadth-First Search 👍 2312 👎 78

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	queue := list.New()
	queue.PushBack(root)
	depth := 0
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}
		}
		depth++
	}
	return depth
}

//leetcode submit region end(Prohibit modification and deletion)
