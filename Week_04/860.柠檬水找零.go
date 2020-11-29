/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	m := make(map[int]int)

	for _, b := range bills {
		if b == 5 {
			m[5]++
		}
		if b == 10 {
			if m[5] > 0 {
				m[5]--
				m[10]++
			} else {
				return false
			}
		}
		if b == 20 {
			if m[10] > 0 && m[5] > 0 {
				m[20]++
				m[10]--
				m[5]--
			} else if m[5] >= 3 {
				m[20]++
				m[5] = m[5] - 3
			} else {
				return false
			}
		}
	}

	return true
}

// @lc code=end

