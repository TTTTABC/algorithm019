/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] 分发饼干
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++ // 判断饼干是否能让孩子满足, 不能就一直等大的饼干
		}
		j++ // 饼干一直在换大的
	}
	return i
}

// @lc code=end

