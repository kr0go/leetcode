package leetcode

/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	tmp := make([]int, m)
	copy(tmp, nums1[:m])
	i, j, k := 0, 0, 0
	for i < m || j < n {
		if i == m {
			nums1[k] = nums2[j]
			j++
			k++
			continue
		}
		if j == n {
			nums1[k] = tmp[i]
			i++
			k++
			continue
		}
		if tmp[i] <= nums2[j] {
			nums1[k] = tmp[i]
			i++
			k++
		} else if tmp[i] > nums2[j] {
			nums1[k] = nums2[j]
			j++
			k++
		}
	}
}

// @lc code=end
