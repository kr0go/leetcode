package leetcode

/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除排序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}

	int := 0
	for j := 1; j < len(nums); j++ {
		if nums[int] != nums[j] {
			int++
			nums[int] = nums[j]
		}
	}
	return int + 1
}

//
//func removeDuplicates(nums []int) int {
//	ind := 0
//	l := len(nums)
//
//	for i := 1; i < l; i++ {
//		if nums[ind] == nums[ind+1] {
//			nums = append(nums[:ind+1], nums[ind+2:]...)
//		} else {
//			ind++
//		}
//	}
//	return len(nums)
//}
// @lc code=end
