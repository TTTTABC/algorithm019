/*
 * @lc app=leetcode.cn id=47 lang=golang
 *
 * [47] 全排列 II
 */

// @lc code=start
func permuteUnique(nums []int) [][]int {
	ret := [][]int{}
	var dfs func(int)
	dfs = func(index int) {
		if index == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, nums)
			ret = append(ret, temp)
			return
		}
		m := map[int]int{}
		for i := index; i < len(nums); i++ {
			if _, ok := m[nums[i]]; ok {
				continue
			}
			nums[i], nums[index] = nums[index], nums[i]
			dfs(index + 1)
			nums[i], nums[index] = nums[index], nums[i]
			m[nums[i]] = 1
		}
	}
	dfs(0)
	return ret
}

// @lc code=end

