/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {
	res := [][]int{}
	var dfs func(int, []int)
	dfs = func(index int, nums []int) {
		if len(nums) == k {
			t := make([]int, k)
			copy(t, nums)
			res = append(res, t)
			return
		}
		for i := index; i <= n; i++ {
			dfs(i+1, append(nums, i))
		}
	}

	dfs(1, []int{})
	return res
}

// @lc code=end

