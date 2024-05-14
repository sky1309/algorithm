package l148

/**
https://leetcode.cn/problems/sort-list
148. 排序链表
给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。

示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]

提示：

链表中节点的数目在范围 [0, 5 * 104] 内
-105 <= Node.val <= 105

进阶：你可以在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 通过快慢指针，获取链表的中心节点， 中心节点为slow
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	midNext := slow.Next
	// 从中间断开链表断开链表
	slow.Next = nil

	// 归并排序，排序左边和右边
	l := sortList(head)
	r := sortList(midNext)

	// mergeList是合并两个有序的链表的操作，上面返回的l和r已经是有序的，所以直接合并就行
	return mergeList(l, r)
}

// 合并两个有序链表的操作
func mergeList(l, r *ListNode) *ListNode {
	// 创建头节点
	dummy := &ListNode{}

	curL := l
	curR := r
	cur := dummy

	for curL != nil && curR != nil {
		if curL.Val <= curR.Val {
			cur.Next = curL
			curL = curL.Next
		} else {
			cur.Next = curR
			curR = curR.Next
		}
		cur = cur.Next
	}

	if curL != nil {
		cur.Next = curL
	} else {
		cur.Next = curR
	}

	return dummy.Next
}
