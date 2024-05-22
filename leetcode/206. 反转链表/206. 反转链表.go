package main

/**
206. 反转链表
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：


输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]
示例 2：


输入：head = [1,2]
输出：[2,1]
示例 3：

输入：head = []
输出：[]


提示：

链表中节点的数目范围是 [0, 5000]
-5000 <= Node.val <= 5000


进阶：链表可以选用迭代或递归方式完成反转。你能否用两种方法解决这道题？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	// prev 记录上一个节点，在cur为nil时，prev就是末尾节点
	var prev *ListNode

	// 当前遍历的节点
	cur := head

	for cur != nil {
		// 临时存储当前的下一个节点
		nxt := cur.Next

		// 当前节点的下一个节点为prev
		cur.Next = prev

		// 更新上一个节点为cur
		prev = cur

		// 继续下一个节点
		cur = nxt
	}
	return prev
}
