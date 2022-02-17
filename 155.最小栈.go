package leetcode

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	val  int
	next *MinStack
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	st := Constructor()
	st.val = x
	st.next = this
	this = &st
}

func (this *MinStack) Pop() {
	this = this.next
}

func (this *MinStack) Top() int {
	return this.val
}

func (this *MinStack) GetMin() int {
	head := this
	min := 0
	for this != nil {
		if min > this.val {
			min = this.val
		}
		this = this.next
	}
	this = head
	return min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
