package leetcode

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	r := len(matrix[0])*len(matrix) - 1
	result := searchArr(matrix, target, 0, r)
	return result
}

func searchArr(matrix [][]int, target int, l, r int) bool {
	if l > r {
		return false
	}
	mlen := len(matrix[0])
	var res bool
	if l <= r {
		h := (l + r) / 2
		if matrix[h/mlen][h%mlen] == target {
			return true
		} else if matrix[h/mlen][h%mlen] < target {
			res = searchArr(matrix, target, h+1, r)
		} else {
			res = searchArr(matrix, target, l, h-1)
		}
	}
	return res
}

// @lc code=end
