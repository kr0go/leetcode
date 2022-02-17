package leetcode

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l *ListNode = &ListNode{}
	pre := l
	carry := 0
	for l1 != nil || l2 != nil {
		pre.Next = &ListNode{}
		p := pre.Next
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		p.Val = (x + y + carry) % 10
		carry = (x + y + carry) / 10
		pre = p

	}
	if carry != 0 {
		pre.Next = &ListNode{Val: carry}
	}
	return l.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//    var l3,l3head *ListNode
//	i:=0
//	for l1!=nil && l2!=nil {
//        tmp,j:=tenProc(l1.Val+l2.Val,i)
//        i=j
//
//		if l3==nil {
//			l3=&ListNode{tmp,nil}
//			l3head=l3
//		} else {
//			l3.Next=&ListNode{tmp,nil}
//			l3=l3.Next
//		}
//
//		l1,l2=l1.Next,l2.Next
//	}
//
//    switch {
//    case l1==nil && l2==nil:
//        if i==1 {
//            l3.Next=&ListNode{i,nil}
//        }
//    case l1==nil && l2!=nil:
//        secondProc(l3,l2,i)
//    case l1!=nil && l2==nil:
//        secondProc(l3,l1,i)
//    default:
//        return l3head
//    }
//    return l3head
//}
//
//func tenProc(l int,i int) (int,int) {
//    if l+i>=10 {
//        return l+i-10,1
//    } else {
//        return l+i,0
//    }
//}
//
//func secondProc(newl *ListNode,l *ListNode,i int){
//    for l!=nil {
//        if i==0 {
//            newl.Next=l
//            break
//        }
//        tmp,j:=tenProc(l.Val,i)
//        newl.Next=&ListNode{tmp,nil}
//        newl=newl.Next
//        i=j
//        if l.Next!=nil {
//            l=l.Next
//        } else if l.Next==nil && i==1 {
//            newl.Next=&ListNode{i,nil}
//            l=nil
//        } else if l.Next==nil && i==0 {
//            l=nil
//        }
//    }
//}

// @lc code=end
