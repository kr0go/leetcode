package leetcode

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var node *ListNode
	var last *ListNode
	for head != nil {
		node = &ListNode{head.Val, last}
		last = node
		head = head.Next
	}
	return node
}

// @lc code=end
