/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) int {
	maxValue := 0
	peak := prices[0]
	valley := prices[0]
	i := 0
	for i < len(prices)-1 {
		for i < len(prices)-1 && prices[i] > prices[i+1] {
			i++
		}
		valley = prices[i]
		for i < len(prices)-1 && prices[i] <= prices[i+1] {
			i++
		}
		peak = prices[i]
		maxValue += peak - valley
	}
	return maxValue
}

// @lc code=end

