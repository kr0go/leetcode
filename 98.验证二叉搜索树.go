package leetcode

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isBST(root, math.MinInt64, math.MaxInt64)
}

func isBST(root *TreeNode, l, r int) bool {
	if root == nil {
		return true
	}

	if root.Val <= l || root.Val >= r {
		return false
	}
	return isBST(root.Left, l, root.Val) && isBST(root.Right, root.Val, r)
}

// @lc code=end
