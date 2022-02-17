package leetcode

/*
 * @lc app=leetcode.cn id=142 lang=golang
 *
 * [142] 环形链表 II
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{}
	flag := 0
	m := make(map[*ListNode]bool)
	for head != nil {
		m[head] = true
		head = head.Next
		if _, ok := m[head]; ok {
			res = head
			flag = 1
			break
		}
	}

	if flag == 0 {
		return nil
	}
	return res
}

// @lc code=end
