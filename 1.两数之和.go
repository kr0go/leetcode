package leetcode

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	res := make([]int, 2)
	for k := range nums {
		if _, ok := m[target-nums[k]]; ok {
			res[0] = k
			res[1] = m[target-nums[k]]
		}
		m[nums[k]] = k
	}
	return res
}

// @lc code=end
