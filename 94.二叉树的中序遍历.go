package leetcode

/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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

//func inorderTraversal(root *TreeNode) []int {
//	if root == nil {
//		return nil
//	}
//	r1 := inorderTraversal(root.Left)
//	r2 := inorderTraversal(root.Right)
//	return append(append(r1, root.Val), r2...)
//}

func inorderTraversal(root *TreeNode) []int {
	st := make([]*TreeNode, 0)
	result := make([]int, 0)
	if root == nil {
		return result
	}

	for root != nil || len(st) != 0 {
		if root != nil {
			st = append(st, root)
			root = root.Left
		} else {
			root = st[len(st)-1]
			result = append(result, root.Val)
			st = st[:len(st)-1]
			root = root.Right
		}
	}
	return result
}

// @lc code=end
