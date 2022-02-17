package leetcode

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	// 只要是负数，肯定不是回文数
	// 若对十求余数为0，表明最后一位是0，肯定不是回文数
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	res := 0
	// 只要反转一半数字再判断是否相等，就知道是不是回文数
	// 而当x小于保存的数时，即表示已经到达中间了
	// e.g. x:12321 res:0 -> x:1232 res:1 -> x:123 res:12 -> x:12 res:123
	// 此时x<=res则跳出循环去判断，分奇数和偶数两种情况，奇数时res要整除10
	for x > res {
		tmp := x % 10
		x /= 10
		res = res*10 + tmp
	}
	return x == res || x == res/10
}

//func isPalindrome(x int) bool {
//
//	if x < 0 {
//		return false
//	}
//
//	i := 0
//	t := x
//	for x != 0 {
//		i = i*10 + x%10
//		x = x / 10
//	}
//	if t == i {
//		return true
//	} else {
//		return false
//	}
//}
// @lc code=end
