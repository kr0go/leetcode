package leetcode

/*
 * @lc app=leetcode.cn id=7 lang=golang
 *
 * [7] 整数反转
 */

// @lc code=start
func reverse(x int) int {
	y := int32(x)
	var z int32
	for y != 0 {
		t := z*10 + y%10
		if t/10 != z {
			return 0
		}
		z = t
		y = y / 10
	}
	return int(z)
}

//func reverse(x int) int {
//	a := 0
//	result := 0
//	for x != 0 {
//		a = x % 10
//		x = x / 10
//		result = result*10 + a
//	}
//	if result > 2147483647 || result < -2147483648 {
//		return 0
//	}
//	return result
//}

// @lc code=end
