package main

/**
234. 回文链表
给你一个单链表的头节点 head ，请你判断该链表是否为
回文链表
。如果是，返回 true ；否则，返回 false 。

示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false


提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9


进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	// 通过快慢指针的方式获取链表的中间节点，slow就是中间节点
	slow := head
	fast := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 反转中间节点后的链表，head链表的尾部节点就是第二段的起始部分
	secondHead := reverseList(slow.Next)

	// 双指针，遍历链表，判断值是否相等
	firstPos := head
	secondPos := secondHead
	for secondPos != nil {
		if firstPos.Val != secondPos.Val {
			return false
		}

		firstPos = firstPos.Next
		secondPos = secondPos.Next
	}

	// 恢复现场
	slow.Next = reverseList(secondHead)

	return true
}

// 反转链表并返回头节点
func reverseList(node *ListNode) *ListNode {
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
