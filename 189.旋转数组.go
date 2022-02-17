package leetcode

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start

func rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	if k == 0 {
		return
	}
	tmp := append(nums[l-k:], nums[:l-k]...)
	copy(nums, tmp)
}

// @lc code=end
