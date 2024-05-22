package l108

/**
108. 将有序数组转换为二叉搜索树
给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵
平衡二叉搜索树。

示例 1：
输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：
输出：[3,1]
解释：[1,null,3] 和 [3,1] 都是高度平衡二叉搜索树。


提示：

1 <= nums.length <= 104
-104 <= nums[i] <= 104
nums 按 严格递增 顺序排列
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func sortedArrayToBST(nums []int) *TreeNode {
//     n := len(nums)
//     if n == 0 {
//         return nil
//     }

//     mid := n/2
//     root := &TreeNode{Val: nums[mid]}
//     root.Left = sortedArrayToBST(nums[:mid])
//     root.Right = sortedArrayToBST(nums[mid+1:])

//     return root
// }

func sortedArrayToBST(nums []int) *TreeNode {
	var dfs func(l, r int) *TreeNode
	dfs = func(l, r int) *TreeNode {
		if l > r {
			return nil
		}

		// 取中间节点作为当前的根节点
		mid := (r-l)/2 + l

		node := &TreeNode{Val: nums[mid]}
		node.Left = dfs(l, mid-1)
		node.Right = dfs(mid+1, r)
		return node
	}

	return dfs(0, len(nums)-1)
}
