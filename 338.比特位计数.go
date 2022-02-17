package leetcode

/*
 * @lc app=leetcode.cn id=338 lang=golang
 *
 * [338] 比特位计数
 */

// @lc code=start

func countBits(num int) []int {
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = dp[i/2]
		}
	}
	return dp
}

//func countBits(num int) []int {
//	res := make([]int, num+1)
//	for i := 0; i <= num; i++ {
//		count := 0
//		tmp := i
//		for j := 0; j <= i; j++ {
//			count += tmp & 1
//			tmp >>= 1
//			if j == i {
//				res[i] = count
//			}
//		}
//	}
//	return res
//}

// @lc code=end
