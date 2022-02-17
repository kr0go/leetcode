package leetcode

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	backtrack(&result, "", 0, 0, n)
	return result
}

func backtrack(str *[]string, s string, left int, right int, n int) {
	if len(s) == 2*n {
		*str = append(*str, s)
		// result = *str
		return
	}

	if left < n {
		backtrack(str, s+"(", left+1, right, n)
	}

	if right < left {
		backtrack(str, s+")", left, right+1, n)
	}
}

// @lc code=end
