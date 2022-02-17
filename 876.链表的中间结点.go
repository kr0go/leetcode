package leetcode

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	s := make([]*ListNode, 0)
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	return s[len(s)/2]
}

// @lc code=end
