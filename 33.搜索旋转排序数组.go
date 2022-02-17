package leetcode

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	ind := findt(nums)
	var result int
	if ind == 0 {
		result = twosep(0, len(nums)-1, nums, target)
	} else if target < nums[0] {
		result = twosep(ind, len(nums)-1, nums, target)
	} else {
		result = twosep(0, ind-1, nums, target)
	}

	return result
}

func twosep(l, r int, nums []int, t int) int {
	if l >= len(nums) || r >= len(nums) {
		return -1
	}

	if l > r {
		return -1
	}
	var ind int
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == t {
			return mid
		} else if nums[mid] < t {
			l = mid + 1
		} else {
			r = mid - 1
		}
		ind = twosep(l, r, nums, t)
	}
	return ind
}

func findt(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i + 1
		}
	}
	return 0
}

// @lc code=end
