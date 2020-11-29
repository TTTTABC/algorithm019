/*
 * @lc app=leetcode.cn id=200 lang=golang
 *
 * [200] 岛屿数量
 */

// @lc code=start
func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var ret = 0
	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x < 0 || y < 0 {
			return
		}
		if x >= len(grid) || y >= len(grid[0]) {
			return
		}
		if grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x-1, y)
		dfs(x, y-1)
		dfs(x+1, y)
		dfs(x, y+1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ret++
				dfs(i, j)
			}
		}
	}

	return ret
}

// @lc code=end

