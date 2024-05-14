package l25

/**
https://leetcode.cn/problems/reverse-nodes-in-k-group
25. K 个一组翻转链表
给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。


示例 1：


输入：head = [1,2,3,4,5], k = 2
输出：[2,1,4,3,5]
示例 2：



输入：head = [1,2,3,4,5], k = 3
输出：[3,2,1,4,5]


提示：
链表中的节点数目为 n
1 <= k <= n <= 5000
0 <= Node.val <= 1000


进阶：你可以设计一个只用 O(1) 额外内存空间的算法解决此问题吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	p1 := dummy
	cur := dummy

	for cur != nil {
		for i := 0; i < k; i++ {
			if cur.Next == nil {
				return dummy.Next
			}
			cur = cur.Next
		}
		// 记录下一个节点，并断开
		nxt := cur.Next
		cur.Next = nil

		// 从p1的next开始反转链表，返回值起始就是cur
		reverse(p1.Next)

		p1Nxt := p1.Next
		p1.Next = cur
		p1 = p1Nxt

		cur = p1
		// 连上节点
		cur.Next = nxt
	}

	return dummy.Next
}

// 反转链表
func reverse(node *ListNode) *ListNode {
	var prev *ListNode
	cur := node
	for cur != nil {
		nxt := cur.Next
		cur.Next = prev
		prev = cur
		cur = nxt
	}
	return prev
}

func reverseKGroup1(head *ListNode, k int) *ListNode {

	// 1.先计算出所有节点的数量，在n<k时结束反转
	n := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		n++
	}

	dummy := &ListNode{}
	dummy.Next = head

	p0 := dummy
	for n >= k {
		// 减少k个节点
		n -= k

		// 一边遍历一边反转链表
		var prev *ListNode
		cur := p0.Next
		for i := 0; i < k; i++ {
			nxt := cur.Next
			cur.Next = prev
			prev = cur
			cur = nxt
		}

		// p0.Next 时下一个p0的位置，需要存储一下
		newP0 := p0.Next
		// 将p0.Next指向反转后的头节点
		p0.Next = prev
		// 更新p0为反转后的尾部节点，即老的 p0.Next
		p0 = newP0
		// 将尾部节点连接到cur，连接链表
		p0.Next = cur
	}

	return dummy.Next
}
