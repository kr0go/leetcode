package leetcode

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start

type stackBracket struct {
	b    string
	next *stackBracket
}

func isValid(s string) bool {
	m := make(map[string]string)
	m["("] = ")"
	m["{"] = "}"
	m["["] = "]"
	var tail *stackBracket = &stackBracket{"", nil}
	for _, v := range s {
		if _, ok := m[string(v)]; ok {
			tmp := tail.next
			tail.next = &stackBracket{string(v), tmp}
		} else if tail.next != nil {
			tail = tail.next
			if m[tail.b] == string(v) {
				continue
			} else {
				return false
			}
		} else {
			return false
		}
	}
	if tail.next != nil {
		return false
	}
	return true
}

// @lc code=end
