/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]
	sum := 0
	for _, v := range nums {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if sum > res {
			res = sum
		}
	}
	return res
}

// @lc code=end
