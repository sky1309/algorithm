package main

/*
101. 对称二叉树
给你一个二叉树的根节点 root ， 检查它是否轴对称。


示例 1：


输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：


输入：root = [1,2,2,null,3,null,3]
输出：false


提示：

树中节点数目在范围 [1, 1000] 内
-100 <= Node.val <= 100


进阶：你可以运用递归和迭代两种方法解决这个问题吗？
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric1(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 迭代的方式
	q := []*TreeNode{root}
	for {
		n := len(q)
		if n == 0 {
			break
		}

		for i := 0; i < n; i++ {
			if q[i] == nil {
				continue
			}
			q = append(q, q[i].Left)
			q = append(q, q[i].Right)
		}

		l, r := 0, n-1
		for l < r {
			if (q[l] == nil && q[r] != nil) || (q[l] != nil && q[r] == nil) {
				return false
			}

			if q[l] != nil && q[r] != nil && q[l].Val != q[r].Val {
				return false
			}

			l++
			r--
		}
		q = q[n:]
	}

	return true
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 递归的方式
	var dfs func(l, r *TreeNode) bool
	dfs = func(l, r *TreeNode) bool {

		if l == nil && r == nil {
			return true
		}

		if l == nil && r != nil || l != nil && r == nil {
			return false
		}

		if l.Val != r.Val {
			return false
		}

		return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
	}

	return dfs(root.Left, root.Right)
}
