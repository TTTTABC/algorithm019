/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	k = len(nums) - k%len(nums)
	copy(nums, append(nums[k:], nums[0:k]...))
}

// @lc code=end

