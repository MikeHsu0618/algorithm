package main

//Given two integer arrays inorder and postorder where inorder is the inorder
//traversal of a binary tree and postorder is the postorder traversal of the same
//tree, construct and return the binary tree.
//
//
// Example 1:
//
//
//Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
//Output: [3,9,20,null,null,15,7]
//
//
// Example 2:
//
//
//Input: inorder = [-1], postorder = [-1]
//Output: [-1]
//
//
//
// Constraints:
//
//
// 1 <= inorder.length <= 3000
// postorder.length == inorder.length
// -3000 <= inorder[i], postorder[i] <= 3000
// inorder and postorder consist of unique values.
// Each value of postorder also appears in inorder.
// inorder is guaranteed to be the inorder traversal of the tree.
// postorder is guaranteed to be the postorder traversal of the tree.
//
//
// Related Topics Array Hash Table Divide and Conquer Tree Binary Tree 👍 5543 ?
//? 83

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 解法一：遞迴 以左閉右閉原則進行切分
// func buildTree(inorder []int, postorder []int) *TreeNode {
//     hash := make(map[int]int, 0)
//     for i, v := range inorder {
//         hash[v] = i
//     }
//     // rootIdx 表示根節點在後序數組中的索引，l 跟 r 表示在中序數組中的前後切分點
//     var build func(rootIdx, l, r int) *TreeNode
//     build = func(rootIdx, l, r int) *TreeNode {
//         // 說明沒有元素，返回空樹
//         if l > r {
//             return nil
//         }
//         // 只剩唯一一個元素，直接返回
//         if l == r {
//             return &TreeNode{Val : inorder[l]}
//         }
//         rootV := postorder[rootIdx]   // 找到後序數組根節點的值
//         rootIn := hash[rootV]         // 找到根節點對應在中序數組的位置
//         root := &TreeNode{Val: rootV} // 構造根節點
//         // 重建左節點和右節點
//         root.Left = build(rootIdx - (r - rootIn) - 1, l, rootIn - 1)
//         root.Right = build(rootIdx - 1, rootIn + 1, r)
//         return root
//     }
//     return build(len(postorder) - 1, 0, len(inorder) - 1)
// }

func buildTree(inorder []int, postorder []int) *TreeNode {
	inHash := make(map[int]int, 0)
	for i, v := range inorder {
		inHash[v] = i
	}
	var build func(postEnd, inStart, inEnd int) *TreeNode
	build = func(postEnd, inStart, inEnd int) *TreeNode {
		if inStart > inEnd {
			return nil
		}
		if inStart == inEnd {
			return &TreeNode{Val: inorder[inStart]}
		}
		rootVal := postorder[postEnd]
		inIdx := inHash[rootVal]
		root := &TreeNode{Val: rootVal}
		root.Left = build(postEnd-1-(inEnd-inIdx), inStart, inIdx-1)
		root.Right = build(postEnd-1, inIdx+1, inEnd)
		return root
	}
	return build(len(postorder)-1, 0, len(inorder)-1)
}

//leetcode submit region end(Prohibit modification and deletion)
