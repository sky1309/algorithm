package main

/*
https://leetcode.cn/problems/flatten-binary-tree-to-linked-list
114. 二叉树展开为链表

给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。


示例 1：


输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]
示例 2：

输入：root = []
输出：[]
示例 3：

输入：root = [0]
输出：[0]


提示：

树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100


进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten1(root *TreeNode) {
	// 递归的方式来处理，先从最深的最右侧节点开始遍历

	var nxt *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		// 中序遍历的逻辑位置
		dfs(node.Left)
		dfs(node.Left)

		dfs(node.Right)
		dfs(node.Left)

		node.Left = nil
		node.Right = nxt
		nxt = node
	}

	dfs(root)
}

func flatten(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			// 如果左子树为空，则直接切换到右子树继续操作
			root = root.Right
		} else {

			pre := root.Left
			// 找到左子树的最右侧的节点，
			for pre.Right != nil {
				pre = pre.Right
			}

			// 将root的有节点接到最右侧的节点上
			pre.Right = root.Right

			// 更新当前root左节点更新到右节点上，并且置空左节点
			root.Right = root.Left
			root.Left = nil

			// 考虑下一个节点
			root = root.Right
		}
	}
}
