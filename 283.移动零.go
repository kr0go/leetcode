package leetcode

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start

func moveZeroes(nums []int) {
	l := len(nums)
	if l < 2 {
		return
	}
	index := 0
	for k := range nums {
		if nums[k] != 0 {
			nums[index] = nums[k]
			index++
		}
	}
	for i := index; i < l; i++ {
		nums[i] = 0
	}
}

// @lc code=end
