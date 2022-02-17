package leetcode

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start

type stackRain struct {
	n    int
	val  int
	next *stackRain
}

var a stackRain = stackRain{0, 0, nil}

func trap(height []int) int {
	if len(height) < 2 {
		return 0
	}
	sum := 0
	result := 0
	var head *stackRain = &a
	for k, v := range height {
		if v > head.val {
			head, result = popCompute(head, v, k)
			sum += result
		}
		tmp := head
		head = &stackRain{k, v, tmp}
	}
	return sum
}

func popCompute(head *stackRain, v int, k int) (*stackRain, int) {
	result := 0
	for head.val < v && head != &a {
		tmp := head
		head = head.next
		if tmp.next == &a {
			break
		}
		if head.val < v {
			result = result + (k-head.n-1)*(head.val-tmp.val)
		} else {
			result = result + (k-head.n-1)*(v-tmp.val)
		}
	}
	return head, result
}

// @lc code=end
