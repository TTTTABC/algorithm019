/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	ret := [][]int{}
	var used = make([]bool, len(nums))
	var path = make([]int, len(nums))
	var dfs func(int)
	dfs = func(l int) {
		for k := range used {
			if !used[k] {
				if l  == len(nums) - 1  {
                    var tmp = make([]int, len(nums))
                    copy(tmp, path)
                    tmp[l] = nums[k]
                    ret = append(ret, tmp)
                    return 
                }else {
                    used[k] = true 
                    path[l] = nums[k]
                    dfs(l + 1)
                    used[k] = false 
                }
			}
		}
	}
	dfs(0)
	return ret
}

// @lc code=end

