package leetcode

/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
}

// bfs
//func minDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	q := make([]*TreeNode, 0)
//	q = append(q, root)
//	count := 1
//	for len(q) > 0 {
//		l := len(q)
//		for i := 0; i < l; i++ {
//			qtop := q[0]
//			q = q[1:]
//			if qtop.Left == nil && qtop.Right == nil {
//				return count
//			}
//			if qtop.Left != nil {
//				q = append(q, qtop.Left)
//			}
//			if qtop.Right != nil {
//				q = append(q, qtop.Right)
//			}
//		}
//		count++
//	}
//	return count
//}

// @lc code=end
