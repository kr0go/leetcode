package leetcode

/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start

func largestRectangleArea(heights []int) int {
	l := len(heights)
	st := make([]int, l+1)
	st[0] = -1
	maxArea, head := 0, 0
	for i := 0; i < l; i++ {
		for st[head] != -1 && heights[st[head]] > heights[i] {
			tmp := heights[st[head]] * (i - st[head-1] - 1)
			head--
			if tmp > maxArea {
				maxArea = tmp
			}
		}
		head++
		st[head] = i
	}
	for st[head] != -1 {
		tmp := heights[st[head]] * (l - st[head-1] - 1)
		head--
		if tmp > maxArea {
			maxArea = tmp
		}
	}
	return maxArea
}

// @lc code=end
