package leetcode

/*
 * @lc app=leetcode.cn id=509 lang=golang
 *
 * [509] 斐波那契数
 */

// @lc code=start

var m map[int]int

func fib(N int) int {
	if N == 1 || N == 0 {
		return N
	}

	if m == nil {
		m = make(map[int]int)
	}
	if _, ok := m[N]; ok {
		return m[N]
	}

	m[N] = fib(N-1) + fib(N-2)
	return fib(N-1) + fib(N-2)
}

// @lc code=end
