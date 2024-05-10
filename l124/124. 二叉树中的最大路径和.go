package l124

import "math"

/**
https://leetcode.cn/problems/binary-tree-maximum-path-sum
124. 二叉树中的最大路径和
二叉树中的 路径 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。

路径和 是路径中各节点值的总和。

给你一个二叉树的根节点 root ，返回其 最大路径和 。
示例 1：


输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6
示例 2：


输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42


提示：

树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {

	ans := math.MinInt
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0 // 没有节点，和为0
		}

		l := max(0, dfs(node.Left))  // 左节点最大链和
		r := max(0, dfs(node.Right)) // 右键点最大链和

		// 假设当前节点为转折节点，计算一次最大和，每个节点都会计算一次
		ans = max(ans, l+r+node.Val)

		return node.Val + max(l, r) // 选择一条最大的和
	}

	dfs(root)

	return ans
}
