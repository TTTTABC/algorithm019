/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	var ret []int
	ret = append(ret, root.Val)
	for _, child := range root.Children {
		ret = append(ret, preorder(child)...)
	}
	return ret
}

// @lc code=end

