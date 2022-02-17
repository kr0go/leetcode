import "sort"

/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start

func permuteUnique(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Ints(nums)

	var (
		ret      [][]int
		generate func([]int, []int)
	)
	generate = func(nums, rest []int) {
		if len(rest) == 0 {
			ret = append(ret, nums)
			return
		}

		for i := 0; i < len(rest); i++ {
			if i > 0 && rest[i] == rest[i-1] {
				continue
			}

			newNums := make([]int, 0, len(nums)+1)
			newNums = append(newNums, nums...)
			newNums = append(newNums, rest[i])
			newRest := make([]int, 0, len(rest)-1)
			newRest = append(newRest, rest[:i]...)
			newRest = append(newRest, rest[i+1:]...)
			generate(newNums, newRest)
		}
	}
	generate(nil, nums)
	return ret
}

//func permuteUnique(nums []int) [][]int {
//	if nums == nil {
//		return nil
//	}
//	result := make([][]int, 0)
//	permutestackback(&result, nums, 0)
//	return result
//}
//
//func permutestackback(result *[][]int, nums []int, ind int) {
//	if ind == len(nums) {
//		tmp := make([]int, len(nums))
//		copy(tmp, nums)
//		*result = append(*result, tmp)
//		return
//	}
//
//	m := map[int]int{}
//	for i := ind; i < len(nums); i++ {
//		if _, ok := m[nums[i]]; ok {
//			continue
//		}
//		nums[i], nums[ind] = nums[ind], nums[i]
//		permutestackback(result, nums, ind+1)
//		nums[i], nums[ind] = nums[ind], nums[i]
//		m[nums[i]]++
//	}
//}
// @lc code=end
