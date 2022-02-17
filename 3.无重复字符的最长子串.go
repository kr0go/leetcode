/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	tmp := 0
	tc := 0
	var count int
	l := len(s)
	if l == 1 {
		return 1
	}
	i := 0
	for i != l {
		for i = tmp + 1; i < l; i++ {
			for j := tmp; j < i; j++ {
				if s[i] == s[j] {
					count = i - tmp
					tmp = j + 1
					if tc < count {
						tc = count
					}
					break
				}
			}
			if i+1 == l {
				count = i - tmp + 1
			}
			if tc < count {
				tc = count
			}
		}
	}
	return tc
}

// @lc code=end
