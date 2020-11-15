/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	m := make(map[[26]int][]string)
	for _, str := range strs {
		k := strArray(str)
		s, ok := m[k]
		if !ok {
			s = make([]string, 0)
		}
		s = append(s, str)
		m[k] = s
	}

	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res

}

func strArray(s string) [26]int {
	res := [26]int{}
	for _, v := range s {
		res[v-'a']++
	}
	return res
}

// @lc code=end

