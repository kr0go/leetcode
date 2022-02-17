package leetcode

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
var m map[int]int

func climbStairs(n int) int {
	if n == 0 || n == 1 || n == 2 {
		return n
	}

	if m == nil {
		m = make(map[int]int)
	}
	if _, ok := m[n]; ok {
		return m[n]
	}
	m[n] = climbStairs(n-1) + climbStairs(n-2)
	return climbStairs(n-1) + climbStairs(n-2)
}

// @lc code=end
