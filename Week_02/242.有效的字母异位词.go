/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := map[rune]int{}
	for _, v := range s {
		hash[v]++
	}
	for _, v := range t {
		vs, ok := hash[v]
		if !ok {
			return false
		}
		if vs == 1 {
			delete(hash, v)
		} else {
			hash[v]--
		}
	}
	return len(hash) == 0
}

// @lc code=end

