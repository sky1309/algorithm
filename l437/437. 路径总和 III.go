package l437

/*
https://leetcode.cn/problems/path-sum-iii/description/?envType=study-plan-v2&envId=top-100-liked
437. 路径总和 III

给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。

路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。

示例 1：
输入：root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8
输出：3
解释：和等于 8 的路径有 3 条，如图所示。
示例 2：

输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
输出：3


提示:

二叉树的节点个数的范围是 [0,1000]
-109 <= Node.val <= 109
-1000 <= targetSum <= 1000
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 表示以root为根节点的数
func rootSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	ans := 0
	// 如果当前根节点值已经等于targetSum则结果+1
	if root.Val == targetSum {
		ans++
	}

	// 分别再计算以左右节点为根节点的路径数目
	ans += rootSum(root.Left, targetSum-root.Val)
	ans += rootSum(root.Right, targetSum-root.Val)

	return ans
}

// 入口
func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	// 以当前root节点为根的路径数目
	ans := rootSum(root, targetSum)

	// 深度遍历左右节点
	ans += pathSum(root.Left, targetSum)
	ans += pathSum(root.Right, targetSum)

	return ans
}
