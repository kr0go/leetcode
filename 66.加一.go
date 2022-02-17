package leetcode

/*
 * @lc app=leetcode.cn id=66 lang=golang
 *
 * [66] 加一
 */

// @lc code=start
func plusOne(digits []int) []int {
	t := 1
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += t
		t = digits[i] / 10
		digits[i] = digits[i] % 10
		if i == 0 && t == 1 {
			a := []int{1}
			digits = append(a, digits...)
			break
		}
		if t == 0 {
			break
		}
	}
	return digits
}

// @lc code=end
