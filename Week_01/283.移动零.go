/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if nums == nil {
		return
	}
	// 双指针法
	// j -> 慢指针
	var j = 0
	// i -> 快指针
	for i := 0; i < len(nums); i++ {
		// i遇到不为0
		if nums[i] != 0 {
			// i和j交换
			nums[j], nums[i] = nums[i], nums[j]
			// j前移
			j++
		}
	}
}

// @lc code=end

