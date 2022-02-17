package leetcode

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	s := make([]*ListNode, 0)
	tmp := head
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	l := len(s)
	if l-n-1 >= 0 {
		a := s[l-n-1]
		a.Next = a.Next.Next
	} else {
		a := s[0]
		a = a.Next
		return a
	}

	return tmp
}

// @lc code=end
