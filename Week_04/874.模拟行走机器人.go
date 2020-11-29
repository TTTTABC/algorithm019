/*
 * @lc app=leetcode.cn id=874 lang=golang
 *
 * [874] 模拟行走机器人
 */

// @lc code=start
func robotSim(commands []int, obstacles [][]int) int {
	res := float64(0)
	x, y := 0, 0
	stepX := []int{0, 1, 0, -1} // 在x轴上的距离   上右下左
	stepY := []int{1, 0, -1, 0} // 在y轴上的距离   上右下左
	cur := 0                    // 上右下左 当前方向 向北
	m := make(map[string]bool)
	for _, v := range obstacles {
		m[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for i := 0; i < len(commands); i++ {
		switch commands[i] {
		case -1:
			cur = (cur + 1) % 4 // 右转
		case -2:
			cur = (cur + 3) % 4 // 左转
		default:
			for j := 0; j < commands[i]; j++ {
				tmpX, tmpY := x+stepX[cur], y+stepY[cur] // 每次循环走一个格子
				if m[fmt.Sprintf("%d-%d", tmpX, tmpY)] { // 遇到障碍物
					break
				}
				x, y = tmpX, tmpY
				res = math.Max(float64(x*x+y*y), res)
			}
		}
	}
	return int(res)
}

// @lc code=end

