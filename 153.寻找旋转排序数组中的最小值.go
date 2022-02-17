package leetcode

/*
 * @lc app=leetcode.cn id=153 lang=golang
 *
 * [153] 寻找旋转排序数组中的最小值
 */

// @lc code=start
func findMin(nums []int) int {
	r := len(nums) - 1
	if r == 0 {
		return nums[0]
	}
	l := 0
	var result int
	if nums[l] < nums[r] {
		result = nums[l]
	} else {
		for l <= r {
			h := (l + r) / 2
			if nums[h] > nums[h+1] {
				result = nums[h+1]
				break
			} else if nums[h-1] > nums[h] {
				result = nums[h]
				break
			} else if nums[h] > nums[r] {
				l = h + 1
			} else {
				r = h - 1
			}
		}
	}

	return result
}

// @lc code=end
