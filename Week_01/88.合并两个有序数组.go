/*
 * @lc app=leetcode.cn id=88 lang=golang
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	l := m + n - 1
	mm := m - 1
	nn := n - 1
	for mm >= 0 && nn >= 0 {
		if nums1[mm] > nums2[nn] {
			nums1[l] = nums1[mm]
			mm--
		} else {
			nums1[l] = nums2[nn]
			nn--
		}
		l--
	}
	for mm >= 0 {
		nums1[l] = nums1[mm]
		l--
		mm--
	}
	for nn >= 0 {
		nums1[l] = nums2[nn]
		l--
		nn--
	}
}

// @lc code=end

