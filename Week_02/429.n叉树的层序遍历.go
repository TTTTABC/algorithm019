/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

var ret [][]int

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}
	ret = [][]int{}
	dfs(root, 0)
	return ret

}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}
	if len(ret) == level {
		ret = append(ret, []int{})
	}

	ret[level] = append(ret[level], root.Val)
	for _, child := range root.Children {
		dfs(child, level+1)
	}
}

// @lc code=end

